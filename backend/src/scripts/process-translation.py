import os
import json
import shutil

# Define the folder path where the JSON files are stored
folder_path = "util/localeutil/locales"


# 1. Remove all 'hash' keys in translate.*.json files
def remove_hash_keys_from_translate_files() -> None:
    for filename in os.listdir(folder_path):
        if filename.startswith("translate.") and filename.endswith(".json"):
            file_path = os.path.join(folder_path, filename)
            with open(file_path, "r") as f:
                data = json.load(f)

            # Remove the 'hash' key from all keys in the JSON object
            for key in data:
                if "hash" in data[key]:
                    del data[key]["hash"]
                if "one" in data[key]:
                    data[key]["one"] = ""
                if "other" in data[key]:
                    data[key]["other"] = ""

            # Save the modified data back to the file
            with open(file_path, "w") as f:
                json.dump(data, f, indent=4)


# 2. Copy active.en.json to translate.en.json and format it
def copy_and_format_active_en_to_translate_en() -> None:
    active_en_file = os.path.join(folder_path, "active.en.json")
    translate_en_file = os.path.join(folder_path, "translate.en.json")

    # Copy active.en.json to translate.en.json
    shutil.copyfile(active_en_file, translate_en_file)

    # Read the content of translate.en.json
    with open(translate_en_file, "r") as f:
        data = json.load(f)

    # Format the data according to the given rules
    formatted_data = {}
    for key, value in data.items():
        if isinstance(value, str):
            # If the value is a string, format it as {other: "string value"}
            formatted_data[key] = {"other": value}
        elif isinstance(value, dict):
            # If the value is an object, keep it as is
            formatted_data[key] = value

    # Save the formatted data back to translate.en.json
    with open(translate_en_file, "w") as f:
        json.dump(formatted_data, f, indent=4)


# Execute the functions
remove_hash_keys_from_translate_files()
copy_and_format_active_en_to_translate_en()
