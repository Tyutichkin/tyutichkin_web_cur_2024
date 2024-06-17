// для взятия айди из таблицы
let currentGoodId = null;
let currentUserId = null;
let currentStockId = null;
let currentIsPriceDesc = false;
let currentIsCountDesc = false;

function addGood() {
    document.getElementById("add-goods-modal").classList.add("open");
}

async function addGoodModal() {
    const name = document.getElementById("good-name").value;
    const description = document.getElementById("good-description").value;
    const price = parseFloat(document.getElementById("good-price").value);

    const good = {
        name: name,
        description: description,
        price: price,
    };

    try {
        const response = await authenticatedFetch(
            "http://localhost:8083/api/v1/good",
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(good),
            }
        );

        if (response.ok) {
            // Assuming you have a function to reload goods after adding a new one
            loadGoods();
            closeAddGoodModal();
        } else {
            console.error("Failed to add good:", response.statusText);
        }
    } catch (error) {
        console.error("Error adding good:", error);
    }
}

async function addStockModal() {
    const address = document.getElementById("add-stock-address").value;

    const stock = {
        address: address,
    };

    try {
        const response = await authenticatedFetch(
            "http://localhost:8083/api/v1/stock",
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(stock),
            }
        );

        if (response.ok) {
            loadStocks();
            closeAddStockModal();
        } else {
            console.error("Failed to add stock:", response.statusText);
        }
    } catch (error) {
        console.error("Error adding stock:", error);
    }
}

function closeAddGoodModal() {
    document.getElementById("add-goods-modal").classList.remove("open");
}

function editGood(goodId) {
    currentGoodId = goodId;
    document.getElementById("edit-goods-modal").classList.add("open");
}

function closeEditGoodModal() {
    document.getElementById("edit-goods-modal").classList.remove("open");
}

function editGoodForUser(goodId, stockId) {
    if (isNaN(parseInt(goodId)) || isNaN(parseInt(stockId))) {
        alert("Good havent stock id");
        return;
    }

    currentGoodId = parseInt(goodId);
    currentStockId = parseInt(stockId);
    document.getElementById("edit-goods-modal-for-user").classList.add("open");
}

function closeEditGoodForUserModal() {
    document
        .getElementById("edit-goods-modal-for-user")
        .classList.remove("open");
}

async function editGoodModal() {
    const name = document.getElementById("edit-good-name").value;
    const description = document.getElementById("edit-good-description").value;
    const price = parseFloat(document.getElementById("edit-good-price").value);

    const good = {
        id: currentGoodId,
        name: name,
        description: description,
        price: price,
    };

    try {
        const response = await authenticatedFetch(
            "http://localhost:8083/api/v1/good",
            {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(good),
            }
        );

        if (response.ok) {
            // Assuming you have a function to reload goods after editing one
            loadGoods();
            closeEditGoodModal();
        } else {
            console.error("Failed to update good:", response.statusText);
        }
    } catch (error) {
        console.error("Error updating good:", error);
    }
}

function addUser() {
    document.getElementById("add-user-modal").classList.add("open");
}

function closeAddUserModal() {
    document.getElementById("add-user-modal").classList.remove("open");
}

async function addUserModal() {
    const fullname = document.getElementById("add-user-fullname").value;
    const login = document.getElementById("add-user-login").value;
    const password = document.getElementById("add-user-password").value;
    const isAdmin = document.getElementById("add-user-is-admin").checked;
    const user = {
        full_name: fullname,
        login: login,
        password: password,
        is_admin: isAdmin,
    };

    try {
        const response = await authenticatedFetch(
            "http://localhost:8083/api/v1/user",
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(user),
            }
        );

        if (response.ok) {
            // Assuming you have a function to reload goods after adding a new one
            loadUsers();
            closeAddUserModal();
        } else {
            console.error("Failed to add user:", response.statusText);
        }
    } catch (error) {
        console.error("Error adding user:", error);
    }
}

function editUser(userID) {
    currentUserId = userID;
    document.getElementById("edit-user-modal").classList.add("open");
}

function closeEditUserModal() {
    document.getElementById("edit-user-modal").classList.remove("open");
}

async function editUserModal() {
    const fullname = document.getElementById("edit-user-fullname").value;
    const login = document.getElementById("edit-user-login").value;
    const password = document.getElementById("edit-user-password").value;
    const isAdmin = document.getElementById("edit-user-is-admin").checked;
    const user = {
        id: currentUserId,
        full_name: fullname,
        login: login,
        password: password,
        is_admin: isAdmin,
    };

    try {
        const response = await authenticatedFetch(
            "http://localhost:8083/api/v1/user",
            {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(user),
            }
        );

        if (response.ok) {
            // Assuming you have a function to reload goods after adding a new one
            loadUsers();
            closeEditUserModal();
        } else {
            console.error("Failed to edit user:", response.statusText);
        }
    } catch (error) {
        console.error("Error editing user:", error);
    }
}

async function deleteUser(userID) {
    try {
        const response = await authenticatedFetch(
            `http://localhost:8083/api/v1/user/${userID}`,
            {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json",
                },
            }
        );

        if (response.ok) {
            loadUsers();
        } else {
            console.error("Failed to delete user:", response.statusText);
        }
    } catch (error) {
        console.error("Error deleting user:", error);
    }
}

function addGoodToStock() {
    loadStocksID();
    loadGoodsID();
    document.getElementById("add-good-to-stock-modal").classList.add("open");
}

function closeAddGoodToStockModal() {
    document.getElementById("add-good-to-stock-modal").classList.remove("open");
}

async function addGoodToStockModal() {
    const goodId = parseInt(
        document.getElementById("add-good-to-stock-good-id-select").value
    );
    const stockId = parseInt(
        document.getElementById("add-good-to-stock-stock-id-select").value
    );
    const count = parseInt(
        document.getElementById("add-good-to-stock-count").value
    );

    const data = {
        good_id: goodId,
        stock_id: stockId,
        good_count: count,
    };

    try {
        const response = await authenticatedFetch(
            "http://localhost:8083/api/v1/good_stock",
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data),
            }
        );

        if (response.ok) {
            closeAddGoodToStockModal();
            loadGoods();
        } else {
            console.error("Failed to add good to stock:", response.statusText);
        }
    } catch (error) {
        console.error("Error adding good to stock:", error);
    }
}

function editGoodToStock() {
    loadStocksID();
    loadGoodsID();
    document.getElementById("edit-good-to-stock-modal").classList.add("open");
}

function closeEditGoodToStockModal() {
    document
        .getElementById("edit-good-to-stock-modal")
        .classList.remove("open");
}

async function editGoodToStockModal() {
    const goodId = parseInt(
        document.getElementById("edit-good-to-stock-good-id-select").value
    );
    const stockId = parseInt(
        document.getElementById("edit-good-to-stock-stock-id-select").value
    );
    const count = parseInt(
        document.getElementById("edit-good-to-stock-count").value
    );

    const data = {
        good_id: goodId,
        stock_id: stockId,
        good_count: count,
    };

    try {
        const response = await authenticatedFetch(
            "http://localhost:8083/api/v1/good_stock",
            {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data),
            }
        );

        if (response.ok) {
            closeEditGoodToStockModal();
            loadGoods();
        } else {
            console.error("Failed to edit good to stock:", response.statusText);
        }
    } catch (error) {
        console.error("Error editing good to stock:", error);
    }
}

async function editGoodToStockForUserModal() {
    const count = parseInt(document.getElementById("edit-good-price").value);

    const data = {
        good_id: currentGoodId,
        stock_id: currentStockId,
        good_count: count,
    };

    try {
        const response = await authenticatedFetch(
            "http://localhost:8083/api/v1/good_stock",
            {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data),
            }
        );

        if (response.ok) {
            closeEditGoodForUserModal();
            loadGoodsForUser();
        } else {
            console.error("Failed to edit good to stock:", response.statusText);
        }
    } catch (error) {
        console.error("Error editing good to stock:", error);
    }
}

function addStock() {
    document.getElementById("add-stock-modal").classList.add("open");
}

function closeAddStockModal() {
    document.getElementById("add-stock-modal").classList.remove("open");
}

function editStock(stockID) {
    currentStockId = stockID;
    document.getElementById("edit-stock-modal").classList.add("open");
}

function closeEditStockModal() {
    document.getElementById("edit-stock-modal").classList.remove("open");
}

async function editStockModal() {
    const address = document.getElementById("edit-stock-address").value;

    const stock = {
        id: currentStockId,
        address: address,
    };

    try {
        const response = await authenticatedFetch(
            "http://localhost:8083/api/v1/stock",
            {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(stock),
            }
        );

        if (response.ok) {
            // Assuming you have a function to reload goods after editing one
            loadStocks();
            closeEditStockModal();
        } else {
            console.error("Failed to edit stock:", response.statusText);
        }
    } catch (error) {
        console.error("Error editing stock:", error);
    }
}

async function deleteStock(stockID) {
    try {
        const response = await authenticatedFetch(
            `http://localhost:8083/api/v1/stock/${stockID}`,
            {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json",
                },
            }
        );

        if (response.ok) {
            loadStocks();
        } else {
            console.error("Failed to delete stock:", response.statusText);
        }
    } catch (error) {
        console.error("Error deleting stock:", error);
    }
}

function sortGoodsByPrice() {
    currentIsPriceDesc = !currentIsPriceDesc;
    searchGoods();
}

function sortGoodsByCount() {
    currentIsCountDesc = !currentIsCountDesc;
    searchGoods();
}

async function deleteGood(goodID) {
    try {
        const response = await authenticatedFetch(
            `http://localhost:8083/api/v1/good/${goodID}`,
            {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json",
                },
            }
        );

        if (response.ok) {
            loadGoods();
        } else {
            console.error("Failed to delete good:", response.statusText);
        }
    } catch (error) {
        console.error("Error deleting good:", error);
    }
}

async function loadGoods() {
    try {
        const response = await authenticatedFetch(
            "http://localhost:8083/api/v1/good/all"
        );

        const data = await response.json();

        if (response.ok) {
            const goods = data.data.map((good) => ({
                id: good.id,
                name: good.name,
                desc: good.description,
                stockNumber: good.stock_id || "-",
                addedBy: good.created_by_user_full_name || "-",
                price: good.price,
                quantity: good.count,
            }));

            const goodsTable = document.getElementById("goods-table");
            goodsTable.innerHTML = goods
                .map(
                    (good) => `
                <tr>
                    <td>${good.id}</td>
                    <td>${good.name}</td>
                    <td>${good.desc}</td>
                    <td>${good.stockNumber}</td>
                    <td>${good.addedBy}</td>
                    <td>${good.price}</td>
                    <td>${good.quantity}</td>
                    <td class="action-btns">
                        <button class="edit-btn" onclick="editGood(${good.id})">✏️</button>
                        <button class="delete-btn" onclick="deleteGood(${good.id})">🗑</button>
                        <button class="download-btn" onclick="downloadGood(${good.id})">⬇️</button>
                    </td>
                </tr>
            `
                )
                .join("");
        } else {
            console.error("Failed to load goods:", data);
        }
    } catch (error) {
        console.error("Error loading goods:", error);
    }
}

async function loadGoodsForUser() {
    try {
        const response = await authenticatedFetch(
            "http://localhost:8083/api/v1/good/all"
        );
        const data = await response.json();

        if (response.ok) {
            const goods = data.data.map((good) => ({
                id: good.id,
                name: good.name,
                desc: good.description,
                stockNumber: good.stock_id || "-",
                addedBy: good.created_by_user_full_name || "-",
                price: good.price,
                quantity: good.count,
            }));

            const goodsTable = document.getElementById("goods-table-for-user");
            goodsTable.innerHTML = goods
                .map(
                    (good) => `
                <tr>
                    <td>${good.id}</td>
                    <td>${good.name}</td>
                    <td>${good.desc}</td>
                    <td>${good.stockNumber}</td>
                    <td>${good.addedBy}</td>
                    <td>${good.price}</td>
                    <td>${good.quantity}</td>
                    <td class="action-btns">
                        <button class="edit-btn" onclick="editGoodForUser(\'${good.id}\', \'${good.stockNumber}\')">✏️</button>
                    </td>
                </tr>
            `
                )
                .join("");
        } else {
            console.error("Failed to load goods:", data);
        }
    } catch (error) {
        console.error("Error loading goods:", error);
    }
}

async function loadUsers() {
    try {
        const response = await authenticatedFetch(
            "http://localhost:8083/api/v1/user/all"
        );
        const data = await response.json();

        if (response.ok) {
            const users = data.data.map((user) => ({
                id: user.id,
                name: user.full_name,
                login: user.login,
            }));

            const usersTable = document.getElementById("users-table");
            usersTable.innerHTML = users
                .map(
                    (user) => `
                <tr>
                    <td>${user.id}</td>
                    <td>${user.name}</td>
                    <td>${user.login}</td>
                    <td class="action-btns">
                        <button class="edit-btn" onclick="editUser(${user.id})">✏️</button>
                        <button class="delete-btn" onclick="deleteUser(${user.id})">🗑</button>
                    </td>
                </tr>
            `
                )
                .join("");
        } else {
            console.error("Failed to load users:", data);
        }
    } catch (error) {
        console.error("Error loading users:", error);
    }
}

async function loadStocks() {
    try {
        const response = await authenticatedFetch(
            "http://localhost:8083/api/v1/stock/all"
        );
        const data = await response.json();

        if (response.ok) {
            const stocks = data.data.map((stock) => ({
                id: stock.id,
                address: stock.address,
            }));

            const stocksTable = document.getElementById("stocks-table");
            stocksTable.innerHTML = stocks
                .map(
                    (stock) => `
                <tr>
                    <td>${stock.id}</td>
                    <td>${stock.address}</td>
                    <td class="action-btns">
                        <button class="edit-btn" onclick="editStock(${stock.id})">✏️</button>
                        <button class="delete-btn" onclick="deleteStock(${stock.id})">🗑</button>
                    </td>
                </tr>
            `
                )
                .join("");
        } else {
            console.error("Failed to load stocks:", data);
        }
    } catch (error) {
        console.error("Error loading stocks:", error);
    }
}

async function loadStocksID() {
    try {
        const response = await authenticatedFetch(
            "http://localhost:8083/api/v1/stock/all"
        );
        const data = await response.json();

        if (response.ok) {
            select = document.getElementById(
                "add-good-to-stock-stock-id-select"
            );
            select.innerHTML = ""; // Clear any existing options

            data.data.forEach((stock) => {
                const option = document.createElement("option");
                option.value = stock.id;
                option.textContent = `id: ${stock.id} | ${stock.address}`;
                select.appendChild(option);
            });
            select = document.getElementById(
                "edit-good-to-stock-stock-id-select"
            );
            select.innerHTML = ""; // Clear any existing options

            data.data.forEach((stock) => {
                const option = document.createElement("option");
                option.value = stock.id;
                option.textContent = `id: ${stock.id} | ${stock.address}`;
                select.appendChild(option);
            });
        } else {
            console.error("Failed to load stocks:", data);
        }
    } catch (error) {
        console.error("Error loading stocks:", error);
    }
}

async function loadGoodsID() {
    try {
        const response = await authenticatedFetch(
            "http://localhost:8083/api/v1/good/all"
        );
        const data = await response.json();

        if (response.ok) {
            select = document.getElementById(
                "add-good-to-stock-good-id-select"
            );
            select.innerHTML = ""; // Clear any existing options

            data.data.forEach((good) => {
                const option = document.createElement("option");
                option.value = good.id;
                option.textContent = `id: ${good.id} | ${good.name}`;
                select.appendChild(option);
            });
            select = document.getElementById(
                "edit-good-to-stock-good-id-select"
            );
            select.innerHTML = ""; // Clear any existing options

            data.data.forEach((good) => {
                const option = document.createElement("option");
                option.value = good.id;
                option.textContent = `id: ${good.id} | ${good.name}`;
                select.appendChild(option);
            });
        } else {
            console.error("Failed to load goods:", data);
        }
    } catch (error) {
        console.error("Error loading goods:", error);
    }
}

async function searchGoods() {
    const minPrice = document
        .getElementById("good-min-price-input")
        .value.trim();
    const maxPrice = document
        .getElementById("good-max-price-input")
        .value.trim();
    const minCount = document
        .getElementById("good-min-count-input")
        .value.trim();
    const maxCount = document
        .getElementById("good-max-count-input")
        .value.trim();
    const nameFilter = document
        .getElementById("good-name-filter-input")
        .value.trim();

    const minPriceValue = minPrice ? parseFloat(minPrice) : 0;
    const maxPriceValue = maxPrice ? parseFloat(maxPrice) : 0;
    const minCountValue = minCount ? parseInt(minCount) : 0;
    const maxCountValue = maxCount ? parseInt(maxCount) : 0;

    const searchData = {
        min_price: minPriceValue,
        max_price: maxPriceValue,
        min_count: minCountValue,
        max_count: maxCountValue,
        name: nameFilter,
        is_price_desc: currentIsPriceDesc,
        is_count_desc: currentIsCountDesc,
    };

    try {
        const response = await authenticatedFetch(
            "http://localhost:8083/api/v1/search/good",
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(searchData),
            }
        );

        if (response.ok) {
            const data = await response.json();
            if (document.getElementById("goods-table")) {
                updateGoodsTable(data.data);
            } else {
                updateGoodsTableForUser(data.data);
            }
        } else {
            console.error("Failed to stock search:", response.statusText);
        }
    } catch (error) {
        console.error("Error stock search:", error);
    }
}

function updateGoodsTable(data) {
    const goodsTable = document.getElementById("goods-table");
    goodsTable.innerHTML = "";

    data.forEach((good) => {
        const row = document.createElement("tr");
        row.innerHTML = `
            <td>${good.id}</td>
            <td>${good.name}</td>
            <td>${good.description}</td>
            <td>${good.stock_id || "-"}</td>
            <td>${good.created_by_user_full_name || "-"}</td>
            <td>${good.price}</td>
            <td>${good.count}</td>
            <td class="action-btns">
                <button class="edit-btn" onclick="editGood(${
                    good.id
                })">✏️</button>
                <button class="delete-btn" onclick="deleteGood(${
                    good.id
                })">🗑</button>
                <button class="download-btn" onclick="downloadGood(${
                    good.id
                })">⬇️</button>
            </td>
        `;
        goodsTable.appendChild(row);
    });
}

function updateGoodsTableForUser(data) {
    const goodsTable = document.getElementById("goods-table-for-user");
    goodsTable.innerHTML = "";

    data.forEach((good) => {
        const row = document.createElement("tr");
        row.innerHTML = `
            <td>${good.id}</td>
            <td>${good.name}</td>
            <td>${good.description}</td>
            <td>${good.stock_id || "-"}</td>
            <td>${good.created_by_user_full_name || "-"}</td>
            <td>${good.price}</td>
            <td>${good.count}</td>
            <td class="action-btns">
                <button class="edit-btn" onclick="editGood(${
                    good.id
                })">✏️</button>
            </td>
        `;
        goodsTable.appendChild(row);
    });
}

async function searchUsers() {
    const fullname = document.getElementById("user-search-fio").value.trim();
    const login = document.getElementById("user-search-login").value.trim();

    const searchData = {
        fullname: fullname,
        login: login,
    };

    try {
        const response = await fetch(
            "http://localhost:8083/api/v1/search/user",
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(searchData),
            }
        );

        if (response.ok) {
            const data = await response.json();
            updateUsersTable(data.data);
        } else {
            console.error("Failed search goods:", response.statusText);
        }
    } catch (error) {
        console.error("Error search goods:", error);
    }
}

function updateUsersTable(data) {
    const userTable = document.getElementById("users-table");
    userTable.innerHTML = "";

    data.forEach((user) => {
        const row = document.createElement("tr");
        row.innerHTML = `
            <td>${user.id}</td>
            <td>${user.full_name}</td>
            <td>${user.login}</td>
            <td class="action-btns">
                <button class="edit-btn" onclick="editUser(${user.id})">✏️</button>
                <button class="delete-btn" onclick="deleteUser(${user.id})">🗑</button>
            </td>
        `;
        userTable.appendChild(row);
    });
}

async function searchStocks() {
    const address = document
        .getElementById("stock-search-address")
        .value.trim();

    // Формируем объект данных для запроса
    const searchData = {
        address: address,
    };

    try {
        const response = await fetch(
            "http://localhost:8083/api/v1/search/stock",
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(searchData),
            }
        );

        if (response.ok) {
            const data = await response.json();
            updateStocksTable(data.data);
        } else {
            console.error("Failed to stock search:", response.statusText);
        }
    } catch (error) {
        console.error("Error stock search:", error);
    }
}

function updateStocksTable(data) {
    const stocksTable = document.getElementById("stocks-table");
    stocksTable.innerHTML = "";

    data.forEach((stock) => {
        const row = document.createElement("tr");
        row.innerHTML = `
        <td>${stock.id}</td>
        <td>${stock.address}</td>
        <td class="action-btns">
            <button class="edit-btn" onclick="editStock(${stock.id})">✏️</button>
            <button class="delete-btn" onclick="deleteStock(${stock.id})">🗑</button>
        </td>
        `;
        stocksTable.appendChild(row);
    });
}

function triggerFileInput() {
    document.getElementById("json-file-input").click();
}

async function importGoodFromJSON(event) {
    const file = event.target.files[0];

    if (file) {
        const formData = new FormData();
        formData.append("file", file);

        try {
            const response = await authenticatedFetch(
                "http://localhost:8083/api/v1/good/upload",
                {
                    method: "POST",
                    body: formData,
                }
            );

            if (response.ok) {
                loadGoods();
                alert("Good was successfully uploaded!");
            } else {
                const errorData = await response.json();
                console.error("Failed import file:", errorData);
                alert("Failed import file: " + errorData.error);
            }
        } catch (error) {
            console.error("Failed request", error);
            alert("Failed request");
        }
    }
}

async function downloadGood(id) {
    try {
        const response = await authenticatedFetch(
            `http://localhost:8083/api/v1/good/download/${id}`
        );
        if (!response.ok) {
            throw new Error(
                `Failed to download file (HTTP status ${response.status})`
            );
        }

        const blob = await response.blob();
        const url = URL.createObjectURL(blob);

        const a = document.createElement("a");
        a.href = url;
        a.download = "good.json"; // Имя файла для скачивания
        a.style.display = "none";
        document.body.appendChild(a);

        a.click();

        // Чистим
        document.body.removeChild(a);
        URL.revokeObjectURL(url);
    } catch (error) {
        console.error("Error downloading file:", error);
    }
}

async function loginUser() {
    const login = document.getElementById("login-input").value;
    const password = document.getElementById("password-input").value;
    const credentials = {
        login: login,
        password: password,
    };

    console.error(JSON.stringify(credentials));

    try {
        const response = await fetch(
            "http://localhost:8083/api/v1/user/login",
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(credentials),
            }
        );

        if (response.ok) {
            const data = await response.json();
            localStorage.setItem("token", data.token);
            localStorage.setItem("role", data.isAdmin);

            if (data.isAdmin === true) {
                window.location.href = "http://127.0.0.1:5500/goods_admin.html";
            } else {
                window.location.href = "http://127.0.0.1:5500/goods_user.html";
            }
        } else {
            console.error("Failed to login:", response.statusText);
            alert("Login failed.");
        }
    } catch (error) {
        console.error("Error logging in:", error);
        alert("Login failed.");
    }
}

async function authenticatedFetch(url, options = {}) {
    const token = localStorage.getItem("token");

    if (!token) {
        alert("User not logged in.");
        return;
    }

    const headers = {
        Authorization: `Bearer ${token}`,
        ...options.headers,
    };

    try {
        const response = await fetch(url, { ...options, headers });
        if (response.ok) {
            return response;
        } else {
            throw new Error(`Request failed with status ${response.status}`);
        }
    } catch (error) {
        console.error("Error fetching data:", error);
    }
}

document.addEventListener("DOMContentLoaded", () => {
    document.querySelectorAll(".input-group form").forEach((form) => {
        form.addEventListener("submit", (event) => {
            event.preventDefault();
            searchUsers();
        });
    });
    if (document.getElementById("goods-table-for-user")) {
        loadGoodsForUser();
    }
    if (document.getElementById("goods-table")) {
        loadGoods();
    }
    if (document.getElementById("users-table")) {
        loadUsers();
    }
    if (document.getElementById("stocks-table")) {
        loadStocks();
    }
});
