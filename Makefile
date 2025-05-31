.PHONY: auth-secret

# Generate a new AUTH_SECRET for JWT authentication and save it to .env file
# Usage: make auth-secret
# This will overwrite the existing AUTH_SECRET in the .env file
auth-secret:
	@sed -i '/^AUTH_SECRET=/d' .env
	@echo "" >> .env
	@echo "# Generated AUTH_SECRET for JWT authentication secret" >> .env
	@echo "AUTH_SECRET=\"$(shell openssl rand -base64 32 | tr -d '\n')\"" >> .env
	@echo "AUTH_SECRET generated and saved to .env"