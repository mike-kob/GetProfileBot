# GetProfile Telegram Bot
This is code for [GetMe Bot](https://t.me/get_my_profile_bot).
This bot returns profile info that Telegram sends to bot
with every update without any additional permissions.

It may be useful, if you need to look up your telegram ID, for example
for bot development purposes.



## Project structure
- `main.go` - file for debugging purposes locally. You need to
set `TOKEN` env variable and change `YOUR_URL` to the URL, to
which you want to recieve updates (for example, `ngrok` tunnel URL).
- `api/update.go` - file that is used by Vercel to make cloud function.
 Contains handler function ([docs](https://vercel.com/docs/runtimes#official-runtimes/go))
 and the main logic of the bot.  