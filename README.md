# Gator

A multi-player command line tool for aggregating RSS feeds and viewing the posts.

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local postgres database. You can then install gator with the following command:

```bash
go install github.com/linus5304/gator
```

## Config

Create a `.gatorconfig.json` file in your home directory with the following contents:

```json
{
  "db_url": "postgres://user:password@localhost:5432/database?sslmode=disable"
}
```

Replace the values with your postgres database connection string.

## Usage

- `gator register <name>` - Create a new user
- `gator addfeed <url>` - Add a new feed
- `gator agg 30s` - Start the aggregator with a 30s refresh rate
- `gator browse` - View the posts

There are a few other commands you'll need as well:

- `gator login <name>` - Login as a user that already exists
- `gator users` - List all users
- `gator feeds` - List all feeds
- `gator follow <url>` - Follow a feed that already exists in the database
- `gator unfollow <url>` - Unfollow a feed that already exists in the database
- `gator following` - List all feeds a user is following
- `gator reset` - This is to reset your database - [use in development]
