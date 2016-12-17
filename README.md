## Synopsis

Personal RSS reader

## Installation

Provide code examples and explanations of how to get the project.

Create cfg.json in project root dir.  

For example:

```
{  
    "driver": "mysql",  
    "connection_string": "root:root@/gorss?charset=utf8&parseTime=True&loc=Local",  
    "opml_path": "/path/for/save/opml/file",  
    "update_minutes": 30,  
    "page_size": 20,  
    "db_backup_path": "/db/backup/dir",  
    "port": 1111  
}
```

$ cd <project/root/directory>  
$ npm install -g bower gulp && npm install  
$ bower install  
$ gulp dist  
$ go build  
$ ./RssReader  

## License
