<?php
// Database configuration
$servername = "db";
$username = "admin";
$password = "admin";
$dbname = "govsphp";

// Create connection
$conn = new mysqli($servername, $username, $password, $dbname);

// Check connection
if ($conn->connect_error) {
    die("Connection failed: " . $conn->connect_error);
}

// User class
class User {
    public int $EmpNo;
    public string $FirstName;
    public string $LastName;

    function __construct(int $emp_no, string $first_name, string $last_name) {
        $this->EmpNo = $emp_no;
        $this->FirstName = $first_name;
        $this->LastName = $last_name;
    }
}

// Fetch users from database
function fetchUsers($conn) {
    $sql = "SELECT emp_no, first_name, last_name FROM employees";
    $result = $conn->query($sql);

    $users = array();
    if ($result->num_rows > 0) {
        while($row = $result->fetch_assoc()) {
            $user = new User($row["emp_no"], $row["first_name"], $row["last_name"]);
            array_push($users, $user);
        }
    }
    return $users;
}

// Handle HTML response
function renderHtml($users) {
?>
<!DOCTYPE html>
<html>
<head>
    <title>User List</title>
</head>
<body>
    <h1>User List</h1>
    <table>
        <tr>
            <th>ID</th>
            <th>First Name</th>
            <th>Last Name</th>
        </tr>
        <?php foreach ($users as $user): ?>
        <tr>
            <td><?= $user->EmpNo ?></td>
            <td><?= $user->FirstName ?></td>
            <td><?= $user->LastName ?></td>
        </tr>
        <?php endforeach; ?>
    </table> 
</body>
</html>
<?php
}

// Main logic to handle different responses
$path = parse_url($_SERVER["REQUEST_URI"], PHP_URL_PATH);

if ($path === "/users") {
    $users = fetchUsers($conn);
    renderHtml($users);
} elseif ($path === "/json_users") {
    $users = fetchUsers($conn);
    header('Content-Type: application/json');
    echo json_encode($users);
} else {
    http_response_code(404);
    echo "404 Not Found";
}

// Close connection
$conn->close();
?>
