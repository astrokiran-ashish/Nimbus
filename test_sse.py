import requests

def test_sse(user_id, user_type):
    url = f"http://localhost:4444/api/v1/notification/stream?ID={user_id}&Type={user_type}"
    
    # Create a session to keep the connection open
    with requests.Session() as session:
        with session.get(url, stream=True) as response:
            if response.status_code == 200:
                print("Connected to SSE stream:")
                for line in response.iter_lines():
                    if line:
                        print(line.decode('utf-8'))
            else:
                print(f"Failed to connect: {response.status_code}")

if __name__ == "__main__":
    user_id = "dac852b4-a903-4daa-bb82-d5f31fd1967b"  # Replace with actual user ID
    user_type = "Consultant"  # or "Consultant"
    test_sse(user_id, user_type)
