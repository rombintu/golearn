# !/bin/python3
import json
import sys, os

if __name__ == "__main__":
    version = os.getenv("VERSION")    
    inter_store = {}

    with open("./version.json", "r") as f:
        inter_store = json.loads(f.read()) 

    with open("./version.json", "w") as f:
        inter_store["version"] = version
        f.write(json.dumps(inter_store, indent=4))
