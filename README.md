# goTweetsAnalyze
twitter analyzer written in Go lang, insipired on the Python tweets_analyzer by x0rz https://github.com/x0rz/tweets_analyzer

[under development]

#### Current analysis
- word count
- weekly activity distribution
- daily activity distribution
- devices used

![Alt text](https://raw.githubusercontent.com/arnaucode/goTweetsAnalyze/master/screen3.png "screen")

![Alt text](https://raw.githubusercontent.com/arnaucode/goTweetsAnalyze/master/screen2.png "screen")

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
./goTweetsAnalyze
```
- follow the instructions that appears on the terminal
