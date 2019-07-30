## devup
When developing on GitHub, you must first take several steps. For example, create a local repository in the project folder, commit the source, then create a remote repository, push from the local repository to the remote repository ...
The command line tool runs them with one command.

## Usage
```
Usage: devup [options]

Options: 
  --path [value], -p [value]   Target project path(absolute path). If you don't specify it, target current directory.
  --token [value], -t [value]  Your github account access token. This flag is required.
  --help, -h               show help
  --version, -v            print the version          
```

## Example of use

$ devup --path /Users/yuta4j1/Mywork/test/ --token 1234567890abcdefghijklmnopqrstuvwxyzzzzzz

Project path to be initialized:  /Users/yuta4j1/Mywork//test/
✔ git init
✔ git Add
✔ git Commit
✔ created remote repository
  ID:  111111111
  FullName:  yuta4j1/test
  CreatedAt:  2019-07-30 12:40:21 +0000 UTC
  CloneURL:  https://github.com/yuta4j1/test.git
Please input your GitHub account. // Ask your GitHub username and password.
Username: 
: 
Password: 
✔ created 'master' branch
✔ git Push
Done successfully!🎉

## License
MIT.