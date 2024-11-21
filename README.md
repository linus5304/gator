# Gator

Gator is a command-line RSS reader.

## Installation

```bash
go install github.com/linus5304/gator
```

## Usage

```bash
gator login           # Log in to your Gator account
gator addfeed <url>   # Add a new RSS feed
gator browse          # Browse your feeds
```

### Available Commands

- `gator register` - Create a new Gator account
- `gator login` - Authenticate with your Gator account
- `gator reset` - This is to reset your database - [use in development]
- `gator users` - List all users
- `gator agg` - Aggregate all feeds
- `gator feeds` - List all feeds
- `gator addfeed <url>` - Subscribe to a new RSS feed
- `gator follow <feed_id>` - Follow a feed
- `gator unfollow <feed_id>` - Unfollow a feed
- `gator following` - List all feeds a user is following
- `gator browse` - Interactive mode to read your feeds

## Configuration

To work with gator you need to have a postgres database running.

You can run a postgres database in docker with the following command:

```bash
docker run -d --name gator-postgres -e POSTGRES_PASSWORD=gator -p 5432:5432 postgres
```
