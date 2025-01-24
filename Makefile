# Targets
.PHONY: orch orch-local create-branch create-branch-from-main

# Run the Orchestrator service
orch:
	@echo "Running Orchestrator service on branch $(shell git rev-parse --abbrev-ref HEAD)"
	cd app && go run main.go

# Run the Orchestrator service locally
orch-local:
	@echo "Running Orchestrator service locally on branch $(shell git rev-parse --abbrev-ref HEAD)"
	cd app && PORT=8001 go run main.go

# Create a new branch
create-branch:
	@read -p "Enter the new branch name: " branch_name; \
	if [ -z "$$branch_name" ]; then \
		echo "Branch name is required."; \
		exit 1; \
	fi; \
	git checkout -b $$branch_name

# Create a new branch from main
create-branch-from-main:
	@read -p "Enter the new branch name: " branch_name; \
	if [ -z "$$branch_name" ]; then \
		echo "Branch name is required."; \
		exit 1; \
	fi; \
	git checkout main && git pull origin main && git checkout -b $$branch_name
