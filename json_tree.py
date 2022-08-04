
from argparse import ArgumentError


class JsonNodeBase:
    def __init__(self,val) -> None:
        if type(val)==str:
            self.type = "string"
        elif type(val)==int:
            self.type = "int"
        elif type(val)==bool:
            self.type = "bool"
        elif type(val)==float:
            self.type = "float32"
        else:
            self.type = str(type(val))
        self.val = val
        pass
    
    def merge(self, node):
        if self.type == node.type:
            return self
        elif type(node)==JsonNodeBase:
            return JsonNodeObject({
                self.type+"_raw":self,
                node.type+"_raw":node
            })
        elif type(node)==JsonNodeArray:
            return node.merge(self)
        elif type(node) ==JsonNodeObject:
            return node.merge(self)
        else:
            raise TypeError("invalid type:" +str(type(node)))
    
    def to_str(self)->str:
        if type(self.val) == str:
            return '"%s"'%str(self.val).replace('"','\\"')
        elif type(self.val) ==bool:
            if self.val:
                return "true"
            else:
                return "false"
        elif self.val == None:
            return "null"
        else:
            return str(self.val)


class JsonNodeObject:
    def __init__(self, init_list = {}):
        self.type = "object"
        self.dict = init_list
    
    def add_value(self, key, val):
        t = type(val)
        if t not in [JsonNodeArray, JsonNodeBase, JsonNodeObject]:
                val = JsonNodeBase(val)
        if key not in self.dict.keys():
            self.dict[key]=val
        else:
            self.dict[key] = self.dict[key].merge(val)
                
    def merge(self, target_node):
        t = type(target_node)
        if t == JsonNodeBase:
            t_ = target_node.type+"_raw"
            if t_ not in self.dict.keys():
                self.dict[t_]= target_node
                return self
            elif self.dict[t_].type == target_node.type:
                return self
            else:
                raise ValueError(target_node)
        elif t == JsonNodeArray:
            return target_node.merge(self)
        elif t== JsonNodeObject:
            for k,v in target_node.dict.items():
                if k in self.dict.keys():
                    self.dict[k] = self.dict[k].merge(v)
                else:
                    self.dict[k]=v
            return self

    def to_str(self)->str:
        ret = "{"
        for k,v in self.dict.items():
            ret += '"%s":%s,' % (k, v.to_str())
        if ret[-1]==",":
            ret = ret[:-1]
        return ret+"}"
            
        
class JsonNodeArray:
    def __init__(self) -> None:
        self.type = "array"
        self.node = None
        pass
    
    def add_node(self, new_node)->None:
        if self.node == None:
            self.node = new_node
        else:
            self.node = self.node.merge(new_node)
            
    def merge(self,target_node)->JsonNodeObject:
        t = type(target_node)
        if self.node== None and t in [JsonNodeArray, JsonNodeBase, JsonNodeObject]:
            self.node = target_node
            return self

        if t == JsonNodeArray:
            if target_node.node!=None:
                self.node = self.node.merge(target_node.node)
        elif t == JsonNodeBase:
            self.node = self.node.merge(target_node)
        elif t == JsonNodeObject:
            self.node = self.node.merge(target_node)
        else:
            raise TypeError("invalid type:" +str(type(target_node)))
        return self
    
    def to_str(self)->str:
        return "[%s]"%self.node.to_str()
        
def import_from_json(json_obj):
    t = type(json_obj)
    if t == dict:
        node = JsonNodeObject({})
        for k,v in json_obj.items():
            node.add_value(k,import_from_json(v))
        return node
    elif t == tuple or t==list:
        node = JsonNodeArray()
        for item in json_obj:
            node.add_node(import_from_json(item))
        return node
    else:
        return JsonNodeBase(json_obj)
    
# def to_golang_struct(obj, name):
#     if obj ==None:
#         raise ArgumentError(obj)
#     t = type(obj)
#     addon_obj_list = {}
#     if t== JsonNodeBase:
#         return type(obj.val), None
#     elif t == JsonNodeArray:
#         obj, addon_obj = to_golang_struct(obj.node, name)
#         if addon_obj == None:
#             return {"array":True, "properties":obj}
#         return {"array":True, "type_name":name, "properties":obj }, addon_obj
#     elif t== JsonNodeObject:
#         struct = {}
#         for k,v in obj.dict.items():
#             obj, addon_obj = to_golang_struct(v, k)
#             struct[k]=obj
#             if addon_obj !=None:
#                 addon_obj_list[k] = addon_obj
#         addon_obj_list[name]= struct
#         return struct, addon_obj_list

class ObjectAnalyser:
    def __init__(self) -> None:
        self.queue = []

    def get_basic_type_name(self, t):
        if t==str:
            return "string"
        elif t==int:
            return "int"
        elif t==bool:
            return "bool"
        elif t==float:
            return "float"
        else:
            return str(t)

    def get_type_recursive(self, name, node):
        t = type(node)
        if t == JsonNodeBase:
            return node.type
        elif t == JsonNodeArray:
            return "[]"+self.get_type_recursive(name, node.node)
        elif t == JsonNodeObject:
            self.queue.append((name, node))
            return name

    def analyse(self, root_name, obj):
        ret = {}
        self.queue = []
        t = type(obj)
        if t!= JsonNodeObject:
            print("unsupport type:", t)
            return {}
        self.queue.append((root_name, obj))
        while(len(self.queue)!=0):
            head = self.queue.pop(0)
            key = head[0]
            value = head[1]
            t = type(value)
            if t!= JsonNodeObject:
                print("unsupport type:", t)
                return {}
            
            properties = {}
            for k,v in value.dict.items():
                properties[k] = self.get_type_recursive(key+"_"+k, v)
            ret[key] = properties
        return ret            
                
                
    def generate_struct(self, package_name:str, list:dict)->str:
        FILE_TEMPLATE='''package %s
        
%s
'''
        STRUCT_TEMPLATE='''type %s struct {
%s
}'''
        PROPERTY_TEMPLATE="\t%s %s `json:\"%s,omitempty\"`"
        structs = ""
        for k,v in list.items():
            properties = ""
            for k1, v1 in v.items():
                if k1[0].isdigit():
                    k1 = "_"+k1
                properties ="%s\n%s"%(properties, PROPERTY_TEMPLATE%(k1.capitalize(),v1, k1))
            structs ="%s\n%s"%(structs, STRUCT_TEMPLATE%(k,properties))
        return FILE_TEMPLATE%(package_name, structs)
            
    