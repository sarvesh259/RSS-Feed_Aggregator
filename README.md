# RSS Feed Aggregator Backend Server

The **RSS Feed Aggregator Backend Server** is a Golang-based server application designed to aggregate and manage RSS feeds from various sources. It provides an efficient way to gather, update, and serve RSS feed data to clients.

This project aims to simplify the process of fetching and managing multiple RSS feeds by providing a centralized server. The backend server periodically fetches the RSS feeds, updates the data, and exposes endpoints for clients to retrieve the aggregated feed content.

## Features

- Aggregates multiple RSS feeds in one place.
- Periodically updates and refreshes feed content (every 10 minutes).
- Provides RESTful API endpoints for client applications to fetch aggregated feed data.
- Easy-to-use and extendable codebase.
