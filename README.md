# BACKEND FOR DSD CENSUS

## use

- git clone \<url>
- make the config file conf/conf.yml, there is a example in the conf/ folder
- go run main.go

for now this will only generate a uuid by discord id and query by discord id

## docs

POST /uuid <br>
used to make a new uuid with a discord id and retrive the results

body:
```
discord_id: str
```

response:
```
uuid: int
discord_id: int
```

GET /uuid/{discord_id} <br>
used to retrive the uuid by discord id in the param

response:
```
uuid: int
discord_id: int
```

# TODOs

- connect it to toromi with authentication that only toromi has for the uuid stuff
- endpoints to save and retrieve census results by existing uuids