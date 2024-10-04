# Hotel Correspondence Finder API

This is a Flask-based web service designed to find correspondences between a target hotel and a list of candidate hotels based on their names and addresses. It leverages the FastEmbed library with pre-trained sentence transformers("sentence-transformers/paraphrase-multilingual-MiniLM-L12-v2") to create semantic embeddings and then applies cosine similarity to determine matches.

## Features
- Semantic Matching: Uses sentence embeddings to compare the semantic similarity between hotel names and addresses.
- RESTful API: Provides a simple HTTP interface to query for correspondences.
- Dynamic Threshold: Allows setting a minimum threshold for considering a match valid.

## Running the Service
Run the Flask app using the following command:
```bash
python app.py
```

The service will be accessible at `http://localhost:5005`.

## API Documentation
Endpoint: `/check_correspondences`
Method: `POST`

Description: Accepts a JSON payload containing a target hotel and a list of candidate hotels. It returns the index of the best match among the candidates if the similarity score exceeds a predefined threshold.

### Expected Data Format:

The request body should be a JSON object with the following structure:

```
{
    "target": [
        ["hotel_name", "hotel_address"]
    ],
    "candidates": [
        ["candidate_name_1", "candidate_address_1"],
        ["candidate_name_2", "candidate_address_2"],
        ...
    ]
}
```
Where:

- "target" is a single entry array containing the name and address of the target hotel.
- "candidates" is an array of arrays, each containing the name and address of a candidate hotel.
### Response:

If a valid match is found, the response will contain the index of the corresponding candidate in the JSON payload. If no valid match is found, -1 is returned.

Example response:

```
{
    "result": 6
}
```

### Usage Example
Here is a sample Python script demonstrating how to use the API:

```python
import requests
import json

def post_correspondences(data):
    url = 'http://localhost:5005/check_correspondences'
    headers = {'Content-Type': 'application/json'}
    response = requests.post(url, json=data, headers=headers)
    
    if response.status_code == 200:
        return response.json()
    else:
        raise Exception(f"Request failed with status {response.status_code}")

# Example data
data = {
    "target": [['Ibis', 'Краснодар, улица Володи Головатого, 306']],
    "candidates": [
        ['Театр', 'Краснодар, Центральный внутригородской округ, микрорайон Центральный, улица Будённого, 153/1'],
        ['Hilton Garden Inn Krasnodar', 'Краснодар, Красная улица, 25/2'],
        ['Краснодар Марриотт', 'Краснодар, Красная улица, 120'],
        ['Мартон Пашковский', 'Краснодар, жилой массив Пашковский, улица Гоголя, 42'],
        ['Мартон Милан', 'Краснодар, Фестивальный микрорайон, улица имени Тургенева, 19'],
        ['Мартон Palace', 'Краснодар, Центральный внутригородской округ, микрорайон Центральный, улица Мира, 70'],
        ['Ibis', 'Краснодар, улица Володи Головатого, 306']
    ]
}

correspond_element = post_correspondences(data)
print(correspond_element)
```