# URL Shortener
Created a URL shortener using Go! The underlying data structure storing the URL mapping is a `map[string]string`. The program can also read URL mappings from a YAML file.

The YAML file format should be a list of objects like such:
```
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
```
### Usage
There is one flag that can be configured:
```
--filename <The YAML file containing the URI mappings>
```