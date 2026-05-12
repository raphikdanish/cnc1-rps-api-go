# Rock Paper Scissors API

A cloud-native REST API for playing Rock Paper Scissors with game statistics, built with Go, React, Docker, and Kubernetes.

## Prerequisites

### For Local Development
- Go 1.20+
- Node.js 18+
- MySQL 8.0+ (running on `localhost:3306`)
- Git

### For Kubernetes Deployment
- Docker
- Kubernetes cluster (minikube, Docker Desktop, or cloud provider)
- kubectl CLI

---

## Running Locally

### 1. Start MySQL

Make sure MySQL is running and accessible on `localhost:3306`.

```bash
# Example: Start MySQL using Docker
docker run --name mysql-local -e MYSQL_ROOT_PASSWORD=pass -p 3306:3306 mysql:8
```

**Note:** The app will attempt to auto-create the database table on startup.

### 2. Start the Backend

In the repo root:

```bash
cd c:\Prv\Trainings\CloudNative-yaha\Assignment2\rps-api-go

# Install dependencies
go mod download

# Run the server
go run main.go
```

Expected output:
```
2026/05/12 12:00:00 Loaded .env.local
2026/05/12 12:00:00 DB_HOST: localhost
2026/05/12 12:00:01 Connected to MySQL
2026/05/12 12:00:01 Games table ready
2026/05/12 12:00:01 Listening on :8080
```

The API will be available at `http://localhost:8080`.

### 3. Start the Frontend

In a new terminal:

```bash
cd c:\Prv\Trainings\CloudNative-yaha\Assignment2\rps-api-go\frontend

# Install dependencies
npm install

# Start dev server
npm run dev
```

Expected output:
```
  VITE v5.4.1  ready in 145 ms

  ➜  Local:   http://localhost:5173/
```

### 4. Access the Application

- **Frontend UI:** `http://localhost:5173`
- **API Docs:** `http://localhost:8080/swagger/`
- **Health check:** `http://localhost:8080/health`

### Local Environment File

The app uses `.env.local` for local settings. If not present, edit `.env.local`:

```env
DB_USER=root
DB_PASSWORD=pass
DB_HOST=localhost
DB_PORT=3306
DB_NAME=rpsdb
```

---

## API Endpoints

### POST /play

Submit a player choice and get the game result.

**Request:**
```json
{
  "choice": "rock"
}
```

Valid choices: `rock`, `paper`, `scissors`

**Response:**
```json
{
  "player_choice": "rock",
  "computer_choice": "paper",
  "result": "loss"
}
```

### GET /stats

Get cumulative game statistics.

**Response:**
```json
{
  "games_played": 10,
  "wins": 4,
  "losses": 3,
  "draws": 3,
  "win_percentage": 40.0
}
```

### GET /health

Health check endpoint.

**Response:**
```
OK
```

---

## Running on Kubernetes

### 1. Prepare the Docker Image

Build and push the API image to your Docker registry:

```bash
cd c:\Prv\Trainings\CloudNative-yaha\Assignment2\rps-api-go

# Build the image
docker build -t your-registry/rps-api-go:latest .

# Push to registry
docker push your-registry/rps-api-go:latest
```

**Note:** Update the image name in `k8s/api-deployment.yaml` to match your registry.

### 2. Update Kubernetes Manifests

Edit `k8s/api-deployment.yaml`:
```yaml
containers:
  - name: rps-api
    image: your-registry/rps-api-go:latest  # Update this line
```

The default configurations in `k8s/` use:
- `k8s/configmap.yaml` for environment variables
  - `DB_HOST: mysql-service` (service DNS name within cluster)
- `k8s/mysql-deployment.yaml` for the MySQL database
- `k8s/api-deployment.yaml` for the API service
- `k8s/api-service.yaml` to expose the API
- `k8s/mysql-service.yaml` to expose MySQL internally

### 3. Deploy to Kubernetes

Apply all manifests in order:

```bash
cd k8s

# Apply ConfigMap (environment variables)
kubectl apply -f configmap.yaml

# Apply MySQL (database)
kubectl apply -f mysql-deployment.yaml
kubectl apply -f mysql-service.yaml

# Apply API (backend)
kubectl apply -f api-deployment.yaml
kubectl apply -f api-service.yaml
```

### 4. Verify Deployment

Check pod status:

```bash
kubectl get pods
```

Expected output:
```
NAME                        READY   STATUS    RESTARTS   AGE
rps-api-xxxxx               1/1     Running   0          10s
mysql-xxxxx                 1/1     Running   0          15s
```

Check service endpoints:

```bash
kubectl get svc
```

### 5. Access the API on Kubernetes

Forward the API port locally:

```bash
kubectl port-forward svc/rps-api 8080:8080
```

Then access:
- **API:** `http://localhost:8080`
- **Swagger Docs:** `http://localhost:8080/swagger/`

Or expose via ingress/load balancer depending on your cluster setup.

### 6. Check Logs

View API logs:

```bash
kubectl logs -f deployment/rps-api
```

View MySQL logs:

```bash
kubectl logs -f deployment/mysql
```

---

## Cleanup

### Local

Kill the frontend and backend processes.

```bash
# Stop MySQL container (if running in Docker)
docker stop mysql-local
docker rm mysql-local
```

### Kubernetes

Delete all resources:

```bash
cd k8s

kubectl delete -f api-service.yaml
kubectl delete -f api-deployment.yaml
kubectl delete -f mysql-service.yaml
kubectl delete -f mysql-deployment.yaml
kubectl delete -f configmap.yaml
```

Or delete the entire namespace:

```bash
kubectl delete namespace rps-api  # if deployed in a specific namespace
```

---

## Environment Variables

### Local Development (`.env.local`)
```env
DB_USER=root
DB_PASSWORD=pass
DB_HOST=localhost
DB_PORT=3306
DB_NAME=rpsdb
```

### Kubernetes Deployment (from ConfigMap)
```yaml
DB_USER: rpsuser
DB_PASSWORD: rpspassword
DB_HOST: mysql-service
DB_PORT: "3306"
DB_NAME: rpsdb
```

**Note:** `DB_HOST` differs because:
- **Local:** Use `localhost` to connect to your host machine's MySQL.
- **Kubernetes:** Use `mysql-service` (service DNS name) to connect within the cluster.

---

## Troubleshooting

### "Connection refused" error

**Local:** Make sure MySQL is running on port 3306.

```bash
# Check if MySQL is running
netstat -an | findstr 3306

# Or restart MySQL
docker start mysql-local
```

**Kubernetes:** Check if MySQL pod is running.

```bash
kubectl logs deployment/mysql
kubectl describe pod <mysql-pod-name>
```

### "No such host" error (localhost not resolving)

The app defaults to `localhost` if `DB_HOST` is not set. Ensure `.env.local` is present and properly formatted.

### Swagger docs showing "No operations defined"

Regenerate Swagger docs:

```bash
swag init --parseFuncBody --dir "./,./handlers,./models"
```

---

## Project Structure

```
rps-api-go/
├── main.go                 # Entry point
├── go.mod                  # Go dependencies
├── Dockerfile              # Docker image
├── .env.local              # Local development config (git ignored)
├── .env.example            # Example config
├── k8s/                    # Kubernetes manifests
│   ├── api-deployment.yaml
│   ├── api-service.yaml
│   ├── mysql-deployment.yaml
│   ├── mysql-service.yaml
│   └── configmap.yaml
├── db/
│   └── database.go         # Database connection & setup
├── handlers/
│   ├── play.go             # Play endpoint
│   └── stats.go            # Stats endpoint
├── models/
│   ├── game.go
│   └── stats.go
├── game/
│   └── logic.go            # Game decision logic
├── docs/
│   ├── docs.go             # Swagger spec (generated)
│   ├── swagger.json
│   └── swagger.yaml
└── frontend/               # React UI
    ├── package.json
    ├── vite.config.js
    ├── index.html
    └── src/
        ├── main.jsx
        ├── App.jsx
        └── style.css
```
