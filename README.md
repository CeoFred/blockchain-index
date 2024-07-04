# Blockchain Indexer

Blockchain Indexer is a Go-based application designed to index and store blockchain events into a PostgreSQL database. It monitors specific smart contract events and keeps a record of user actions, allowing for efficient querying and analysis of blockchain activities.

## Table of Contents

- [About the Project](#about-the-project)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## About the Project

The Blockchain Indexer project aims to provide a reliable and efficient way to track blockchain events and store them in a relational database. This enables users to query and analyze blockchain data without needing to interact directly with the blockchain, which can be time-consuming and complex.

## Features

- **Event Logging**: Indexes and stores blockchain events into a PostgreSQL database.
- **User Actions Tracking**: Records user actions derived from blockchain transactions.
- **API Endpoints**: Provides a RESTful API to query stored data.

## Getting Started

To get a local copy up and running, follow these steps.

### Prerequisites

- [Go](https://golang.org/doc/install) 1.16 or higher
- [PostgreSQL](https://www.postgresql.org/download/)
- [Docker](https://www.docker.com/get-started) (optional, for running PostgreSQL in a container)

### Installation

1. **Clone the repository**
   ```sh
   git clone https://github.com/CeoFred/blockchain-indexer.git
   cd blockchain-indexer
   ```

2. **Set up PostgreSQL**
   - **Using Docker:**
     ```sh
     docker run --name blockchain-postgres -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres
     ```
   - **Manual Installation:**
     Follow the instructions on the PostgreSQL website to install it on your machine.

3. **Configure the database connection**
   - Create a `.env` file in the project root:
     ```env
     DB_HOST=localhost
     DB_PORT=5432
     DB_USER=your_db_user
     DB_PASSWORD=your_db_password
     DB_NAME=your_db_name
     RPC_URL=https://enugu-rpc.assetchain.org
     PORT=9900
     ```

4. **Run database migrations**
    Database migrations automatically run when the application is started. migration files are written in SQL and can be found here.

5. **Build and run the project**
  For linux based machines, run the following command
   ```sh
   make run-local
   ```

## Usage

Once the project is running, it will start listening for blockchain events and indexing them into the PostgreSQL database. You can interact with the indexed data via the provided API endpoints.

## API Endpoints

### Add a Contract to Index

To add a contract to the index, use the following API endpoint:

- **URL**: `{{HOST}}/api/v1/contract`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "address": "0x38048BF4Ee7642DB08411c42cA57C3f6bc8F41A7",
    "name": "BEER Token $BEER Token",
    "start_block": 6177,
    "event": [
      {
        "name": "Transfer",
        "description": "token transfer"
      },
      {
        "name": "Approval",
        "description": "Approve token"
      }
    ]
  }
  ```

#### Description

- **address**: The Ethereum address of the smart contract to be indexed.
- **name**: The name of the smart contract.
- **start_block**: The block number from which to start indexing events.
- **event**: An array of events to be indexed, each with a name and description.

#### Example

```bash
curl -X POST {{HOST}}/api/v1/contract \
     -H "Content-Type: application/json" \
     -d '{
           "address": "0x38048BF4Ee7642DB08411c42cA57C3f6bc8F41A7",
           "name": "BEER Token $BEER Token",
           "start_block": 6177,
           "event": [
             {
               "name": "Transfer",
               "description": "token transfer"
             },
             {
               "name": "Approval",
               "description": "Approve token"
             }
           ]
         }'
```


---

### Retrieve Events for a Specific Contract

To retrieve events under a specific contract address, use the following API endpoint:

- **URL**: `{{HOST}}/api/v1/contract-event/:address`
- **Method**: `GET`

#### Response

```json
{
    "data": [
        {
            "id": "019035b7-13d8-72c9-9126-a2b4141893bf",
            "contract_id": "019035b7-13cf-72c7-8c7d-c80231ec20f3",
            "event_id": "019035b7-13d6-72c8-824a-5cbb02b2a797",
            "active": true,
            "contract": {
                "id": "019035b7-13cf-72c7-8c7d-c80231ec20f3",
                "name": "BEER Token $BEER Token",
                "address": "0x38048BF4Ee7642DB08411c42cA57C3f6bc8F41A7",
                "start_block": 6177,
                "end_block": 6762,
                "created_at": "2024-06-20T13:55:09.007064Z",
                "updated_at": "2024-06-21T10:10:54.571941Z",
                "deleted_at": "0001-01-01T00:00:00Z"
            },
            "event": {
                "id": "019035b7-13d6-72c8-824a-5cbb02b2a797",
                "name": "Approval",
                "signature": "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925",
                "description": "Approve token",
                "created_at": "2024-06-20T13:55:09.014246Z"
            },
            "created_at": "2024-06-20T13:55:09.016855Z",
            "deleted_at": "0001-01-01T00:00:00Z"
        }
    ],
    "message": "success",
    "status": true
}
```

---

### Retrieve All Event Logs

To retrieve all event logs for a specific contract address, use the following API endpoint:

- **URL**: `{{HOST}}/logs/:address`
- **Method**: `GET`
- **Query Parameters** (all optional):
  - `event_name`: The name of the event (e.g., `Approval`).
  - `block_number`: The specific block number.
  - `from_block`: The starting block number.
  - `to_block`: The ending block number.

#### Sample Request

```bash
curl http://localhost:9900/logs/0x38048BF4Ee7642DB08411c42cA57C3f6bc8F41A7?event_name=Approval&block_number=6282&from_block=62&to_block=6394
```

#### Response

```json
{
    "data": [
        {
            "id": "01903101-0e9c-77c9-8ce9-e91f8ad1537b",
            "contract_event_id": "01903100-fb60-77c8-a5bb-3527e482cf4f",
            "contract_id": "01903100-fb54-77c7-b938-67a923814d71",
            "contract_address": "0x38048BF4Ee7642DB08411c42cA57C3f6bc8F41A7",
            "event_name": "Transfer",
            "block_number": 6282,
            "transaction_hash": "0xe1ce554fb9309a9fd2e253fea16e853c09ce93840f5850f7b4cf2d08a91a5563",
            "log_index": 0,
            "data": {
                "raw": "0000000000000000000000000000000000000000033b2e3c9fd0803ce8000000",
                "formatted": {
                    "To": "0x5daa96364bd8e0c4f95004adb4bde0f2afe933c9",
                    "From": "0x0000000000000000000000000000000000000000",
                    "Value": 1000000000000000000000000000
                }
            },
            "topics": [
                "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                "0x0000000000000000000000000000000000000000000000000000000000000000",
                "0x0000000000000000000000005daa96364bd8e0c4f95004adb4bde0f2afe933c9"
            ],
            "timestamp": "0001-01-01T00:00:00Z",
            "created_at": "2024-06-19T15:57:51.873085Z"
        }
    ],
    "message": "success",
    "status": true
}
```

#### Description

- **id**: The unique identifier of the event log.
- **contract_event_id**: The ID of the contract event.
- **contract_id**: The ID of the contract.
- **contract_address**: The Ethereum address of the contract.
- **event_name**: The name of the event (e.g., `Transfer`).
- **block_number**: The block number where the event occurred.
- **transaction_hash**: The hash of the transaction that generated the event.
- **log_index**: The index of the log within the block.
- **data**:
  - **raw**: The raw data of the log.
  - **formatted**: A formatted object containing relevant data fields:
    - **To**: The recipient address.
    - **From**: The sender address.
    - **Value**: The value transferred.
- **topics**: An array of topics associated with the event.
- **timestamp**: The timestamp of the event.
- **created_at**: Timestamp of when the log was created.

---


### User Actions (Not implemented)
- **Get user actions**
  - **Endpoint**: `GET /users/:address/actions`
  - **Description**: Retrieve actions performed by a specific user.

### Contract Analytics (Not implemented)
- **Get detailed analytics for contract**
  - **Endpoint**: `GET /analytics/:address`
  - **Description**: 


## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

---