## Usage cli
```bash
$ golearn ping
INFO[0000] {"message":"pong"}

$ golearn auth --help
NAME:
   main auth - Authentification (get token)

USAGE:
   main auth [command options] [arguments...]

OPTIONS:
   --login value  Your login
   --pass value   Your password
   --role value   Your role
   
$ golearn auth --login user1 --pass 123
INFO[0000] Your token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9

```