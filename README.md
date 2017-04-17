# goTweetsAnalyze
twitter analyzer written in Go lang, insipired on the Python tweets_analyzer by x0rz https://github.com/x0rz/tweets_analyzer

#### Argus Panoptes
https://en.wikipedia.org/wiki/Argus_Panoptes

https://en.wikipedia.org/wiki/Panopticon


![Argus](https://raw.githubusercontent.com/arnaucode/argus/master/argus.jpg "argus")

[under development]

#### Current features
- User analyzer
    - word count
    - weekly activity distribution
    - daily activity distribution
    - devices used
- Delete Tweets and Favs

![screen](https://raw.githubusercontent.com/arnaucode/argus/master/screen3.png "screen")

![screen](https://raw.githubusercontent.com/arnaucode/argus/master/screen2.png "screen")

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
./argus
```
- follow the instructions that appears on the terminal
