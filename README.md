# argos
twitter analyzer written in Go lang, insipired on the Python tweets_analyzer by x0rz https://github.com/x0rz/tweets_analyzer

#### Argos Panoptes
https://en.wikipedia.org/wiki/Argus_Panoptes

https://en.wikipedia.org/wiki/Panopticon


![Argos](https://raw.githubusercontent.com/arnaucode/argos/master/argos.jpg "argos")

[under development]

#### Current features
1. User analyzer
    - word count
    - weekly activity distribution
    - daily activity distribution
    - devices used
    - hashtags most used count
    - user mentions coun [not implemented yet]
2. Delete Tweets and Favs
3. Unfollow all [not implemented yet]
4. Random follow [not implemented yet]
    - selects n number of accounts to follow, and follows n random accounts
5. Random tweet [not implemented yet]
    - post a tweet with random content (from newspaper)

![screen](https://raw.githubusercontent.com/arnaucode/argos/master/screen3.png "screen")

![screen](https://raw.githubusercontent.com/arnaucode/argos/master/screen2.png "screen")

needs a twitterConfig.json file on the /build folder with the content:
```
{
    "consumer_key": "xxxxxxxxxxxxxxxx",
    "consumer_secret": "xxxxxxxxxxxxxxxx",
    "access_token_key": "xxxxxxxxxxxxxxxx",
    "access_token_secret": "xxxxxxxxxxxxxxxx"
}

```

to run it:
- go to the /build folder
- open terminal
- execute the script with:
```
./argos
```
- follow the instructions that appears on the terminal
