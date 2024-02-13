### store
The command-line interface (CLI) application developed in Go, to upload, update, delete, and list plain-text files on server ([server's source code](https://github.com/zakisk/redhat-server)). Alongside this, the CLI can output all words in all files on the server and most frequent words of them. The CLI does this in optimal way, taking minimum time to count words and frequent words.


First of all, we need to launch server in docker so that app can make requests to it.

## store deploy

Deploys server to docker

### Synopsis

```
store deploy
```

deploys server on docker
### Options

```
  -h, --help          help for deploy
  -p, --port string   server port (default "9254")
```


## store add
Now, let's upload a file to server
example:
```
store add -f ./data/file.txt
```

output:
```
File `file.txt` is uploaded successfully
```

you can upload multiple files as well

```
store add -f ./data/file1.txt -f ./data/file2.txt
```
### Options
```
  -f, --file stringArray   Files to be uploaded on server
  -h, --help               help for add
```

### SEE ALSO

* [store update](./docs/store_update.md)	 - updates files
* [store rm](./docs/store_rm.md)	 - remove file froms server
* [store ls](./docs/store_ls.md)	 - lists all files of server
* [store wc](./docs/store_wc.md)	 - Counts words in all files on server
* [store freq-words](./docs/store_freq-words.md)	 - Prints N most frequent words on server