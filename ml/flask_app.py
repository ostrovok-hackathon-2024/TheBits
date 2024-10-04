from flask import Flask, request, jsonify
from sklearn.metrics.pairwise import cosine_similarity
import numpy as np
from fastembed import TextEmbedding


app = Flask(__name__)
embedding_model = TextEmbedding("sentence-transformers/paraphrase-multilingual-MiniLM-L12-v2")


@app.route('/hello')
def hello():
    return 'Hello World!'

def find_correspondences(correspndence_check: dict[str, list[list[str]]], min_correspondences_value: float = 0.85) -> int:
    """
    Funtion get dictionary with keys: "target", "candidates".
    Target is containing list where first element the hotel name, the second element is address&
    Candidates is list of list with hotel names and address which we should check.

    Parameters
    ----------
    correspndence_check : dict[str, list[list[str]]]
        Dictionary with keys: "target", "candidates" and values:
        list[list[str]] - list of list with hotel names and address which we should check.

    min_correspondences_value : float
        Minimum value of correspondences.
        
    Returns
    -------
    int
        Return number of correspondences element or -1 if not found
    """

    target = correspndence_check["target"]
    correspondence_candidates = correspndence_check["candidates"]
    
    np_candidates = np.array(correspondence_candidates)
    np_names = np_candidates[:, 0]
    np_addres = np_candidates[:, -1]

    np_target_name = np.array([target[0][0]])
    np_target_addres = np.array([target[0][1]])

    embeddings_names_trg = np.array(list(embedding_model.embed(np_target_name)))
    embeddings_addres_trg = np.array(list(embedding_model.embed(np_target_addres)))

    embeddings_names_correspond = np.array(list(embedding_model.embed(np_names)))
    embeddings_addres_correspond = np.array(list(embedding_model.embed(np_addres)))

    target_concat = np.concatenate((embeddings_names_trg, embeddings_addres_trg), axis=1)
    correspond_concat = np.concatenate((embeddings_names_correspond, embeddings_addres_correspond), axis=1)
    
    similarity = cosine_similarity(target_concat, correspond_concat).squeeze()
    index = similarity.argmax()

    if similarity[index] > min_correspondences_value:
        return index
    else:
        return -1

@app.route('/check_correspondences', methods=['POST'])
def check_correspondences():
    """
    Accepts a JSON payload and returns the index of the best match among the candidates.
    """
    data = request.get_json()

    if not isinstance(data, dict):
        return jsonify({"error": "Invalid data format"}), 400
    
    result = str(find_correspondences(data))
    return jsonify({"result": result})

@app.route('/')
def index():
    return 'Hello. You are arrived at "/"'

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5005)
