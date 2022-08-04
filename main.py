import os
import json

from json_tree import *

t = None
for root, dirs, files in os.walk("./resources"):
    for file in files:
        with open(root+"/"+file, "r", encoding='utf-8') as f:
            obj = json.load(f)
            if t==None:
                t=import_from_json(obj)
            else:
                t.merge(import_from_json(obj))
            
            
txt = t.to_str()
with open("test_raw.json", "w", encoding="utf-8") as f:
    f.write(txt)
    
json_obj = json.loads(txt)
os.makedirs("template", exist_ok=True)
if type(json_obj)==dict:
    for k,v in json_obj.items():
        while type(v) in [tuple, list] and len(v)!=0:  
            v = v[0]
        with open("template/%s.json"%k, "w", encoding="utf-8") as f:
                json.dump(v, f, indent=4, ensure_ascii=False)


#test
TEMPLATE_NAME = "baseitem"
with open("template/%s.json"%TEMPLATE_NAME,"r", encoding='utf-8') as f:
    json_obj = json.load(f)
    t=import_from_json(json_obj)
    conv = ObjectAnalyser()
    ret = conv.analyse(TEMPLATE_NAME, t)
    with open("%s.go"%TEMPLATE_NAME, "w", encoding="utf-8") as f1:
        s = conv.generate_struct(TEMPLATE_NAME, ret)
        f1.write(s)