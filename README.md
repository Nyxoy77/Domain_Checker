```markdown
# Domain Checker

## Overview
**Domain Checker** is a tool designed to help verify the email readiness of a domain by checking essential DNS records, specifically MX, SPF, and DMARC records. With this project, you can ensure a domain's email setup and enhance email deliverability and security.

This project is currently implemented as a backend tool using Go, with plans to add a frontend interface for a more user-friendly experience.

## Features
- **MX Record Check**: Verifies if a domain has an MX (Mail Exchanger) record, which indicates email service capability.
- **SPF Record Check**: Checks for the presence of an SPF (Sender Policy Framework) record, which specifies authorized email servers for the domain.
- **DMARC Record Check**: Looks for a DMARC (Domain-based Message Authentication, Reporting & Conformance) record, which provides additional email security and reporting.

## Future Plans
- Develop a frontend interface to make the tool accessible to non-technical users.
- Add additional checks or reporting features to further enhance the tool’s capabilities.

## Installation
1. Clone this repository:
   ```bash
   git clone https://github.com/Nyxoy77/Domain_Checker.git
   ```
2. Navigate to the project directory:
   ```bash
   cd Domain_Checker
   ```
3. Run the Go program:
   ```bash
   go run main.go
   ```

## Usage
1. After running the program, enter the domain you wish to verify when prompted.
2. The program will output whether the domain has valid MX, SPF, and DMARC records along with the record details if available.

### Example Output
```plaintext
Enter the domain address you want to verify:
example.com
Domain: example.com
 - Has MX record: true
 - Has SPF record: true
   SPF Record: v=spf1 include:_spf.example.com -all
 - Has DMARC record: true
   DMARC Record: v=DMARC1; p=none; rua=mailto:dmarc-reports@example.com
```

## Project Structure
- `main.go`: Contains the main logic for checking MX, SPF, and DMARC records.

## Technologies
- **Language**: Go
- **Networking**: DNS lookups

## Contributing
Contributions are welcome! Feel free to open an issue or submit a pull request if you'd like to add features or improve the existing functionality.

## License
This project is licensed under the MIT License. See [LICENSE](LICENSE) for more details.

## Connect
Follow me on [LinkedIn](https://www.linkedin.com/in/shivam-rai/) for updates and check out my other projects on [GitHub](https://github.com/Nyxoy77).

Stay tuned for updates as I work on the frontend to make this tool even more user-friendly!

---

**Shivam Rai**
```
