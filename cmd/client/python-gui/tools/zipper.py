import zipfile

def Create_zip_file(path, data):
    with zipfile.ZipFile(path, mode="w") as archive:
        for el in data:
            with archive.open(el["title"] + ".pdf", "w") as new_file:
                try:
                    new_file.write(el["content"])
                except Exception as e:
                    print(e)
                    