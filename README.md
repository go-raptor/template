# Raptor Project Template
This is a template for creating new Raptor projects. It includes a basic setup with the HelloController and HelloService.

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/go-raptor/template.git
   cd template
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Install Raptor CLI if not already installed:
   ```bash
   go install github.com/go-raptor/cli/cmd/raptor@latest
   ```

4. Run the application:
   ```bash
   raptor dev
   ```

5. Access the API:
   - http://localhost:3000/api/v1/hello