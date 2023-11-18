# Toped-Scrapper

## Overview
Toped-Scrapper is a Go-based application for scraping product information from Tokopedia. It uses a CLI for user interaction, allowing users to specify search parameters like product category, limit, and number of workers. The application fetches product data asynchronously using goroutines and worker pools, stores the data in a PostgreSQL database, and outputs the results in a CSV file.

## Terminal Input
![Screenshot 2023-11-18 163102](https://github.com/jonatanlaksamanapurnomo/go-scrapper/assets/39803159/3d1819c5-5e47-4e95-a7e4-e28426a5c298)

## Terimnal Output 
![Screenshot 2023-11-18 163124](https://github.com/jonatanlaksamanapurnomo/go-scrapper/assets/39803159/80e091c6-8012-49da-8df2-5dbce9298046)

## Features
- Fetch products from Tokopedia based on user-defined parameters.
- Asynchronous data fetching and insertion using goroutines and worker pools.
- Storing fetched data in PostgreSQL.
- Generating a CSV report of fetched products.
- Dockerized setup for easy deployment and environment management.

## Getting Started
These instructions will guide you through setting up and running Toped-Scrapper on your local machine.

### Prerequisites
- Docker
- Make
- Golang

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/jonatanlaksamanapurnomo/go-scrapper

### Run
1. Go Root directory:
   ```bash
   make go-run-cli
