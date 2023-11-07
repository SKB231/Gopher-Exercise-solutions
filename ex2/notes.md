## Some additonal notes about ex2:

- Had a few issues here with go.mod. But realized later on that to import locally through subdirectories instead of determining package doesn't exist the subdirectory names should also be consistent with the package name of the file in the subdirecory.
- Yaml files have a much different format compared to JSON files such as extra spaces at the beginning causing issues. But other than that, yaml sequences are similar to arrays except they start with a single '-'
