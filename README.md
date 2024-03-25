# RSS Aggregator

## Overview

This project is a personal endeavor focused on learning how to develop an RSS Aggregator while also gaining proficiency in REST API and Go development. The goal is to create a platform where users can subscribe to and manage RSS feeds from various sources, as well as follow their own posts through user accounts.

## Features

- **RSS Feed Subscription:** Users can add and manage RSS feeds from different sources.
- **User Accounts:** Users can create accounts to follow their own posts and manage subscriptions.
- **Feed Updates:** Automatically fetches and updates feeds at regular intervals.
- **Local Database:** Stores user preferences and subscriptions in a PostgreSQL database.

## Technologies Used

- **Go Language:** Backend development using Go.
- **Goose Library & SQLC:** Connects Golang with PostgreSQL for database operations.
- **PostgreSQL:** Database to store user data, feed subscriptions, and posts.

## API Endpoints

### Health Check
- **GET /v1/healthz:** Check if the server is initialized.

### Error Check
- **GET /v1/err:** Test endpoint to receive an error message.

### User Management
- **POST /v1/users:** Create a new user.
- **GET /v1/users:** Login to an existing user account.

### Feed Management
- **POST /v1/feeds:** Create a new feed.
- **GET /v1/feeds:** Get all posted feeds.

### Feed Subscription
- **POST /v1/feed_follows:** Add a new RSS feed subscription for the logged-in user.
- **GET /v1/feed_follows:** Get all RSS feed subscriptions for a specific user.
- **DELETE /v1/feed_follows/{feedFollowID}:** Delete an RSS feed subscription.
