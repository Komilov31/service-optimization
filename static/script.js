const btn = document.getElementById('fetchOrderBtn');
const input = document.getElementById('orderUid');
const display = document.getElementById('orderDisplay');
const errorMsg = document.getElementById('errorMsg');

btn.addEventListener('click', () => {
  const orderUid = input.value.trim();
  if (!orderUid) {
    errorMsg.textContent = 'Пожалуйста, введите order_uid.';
    display.style.display = 'none';
    return;
  }
  errorMsg.textContent = '';
  display.style.display = 'none';
  display.innerHTML = 'Загрузка...';

  const apiUrl = `http://localhost:8081/order/${encodeURIComponent(orderUid)}`;

  fetch(apiUrl)
    .then(response => {
      if (!response.ok) throw new Error('Нету заказа с таким uid');
      return response.json();
    })
    .then(data => {
      display.innerHTML = renderOrder(data);
      display.style.display = 'block';
    })
    .catch(err => {
      display.style.display = 'none';
      errorMsg.textContent = 'Ошибка при получении данных: ' + err.message;
    });
});


function formatDate(isoStr) {
  const d = new Date(isoStr);
  return d.toLocaleString();
}

function formatTimestamp(ts) {
  const d = new Date(ts * 1000);
  return d.toLocaleString();
}

function renderOrder(order) {
  return `
    <h2>Заказ ${order.order_uid}</h2>
    <div class="section">
      <h3>Информация о доставке</h3>
      <div class="field"><strong>Имя:</strong> ${order.delivery.name}</div>
      <div class="field"><strong>Телефон:</strong> ${order.delivery.phone}</div>
      <div class="field"><strong>Email:</strong> ${order.delivery.email}</div>
      <div class="field"><strong>Адрес:</strong> ${order.delivery.region}, ${order.delivery.city}, ${order.delivery.address}, ${order.delivery.zip}</div>
      <div class="field"><strong>Служба доставки:</strong> ${order.delivery_service}</div>
    </div>
    <div class="section">
      <h3>Оплата</h3>
      <div class="field"><strong>Транзакция:</strong> ${order.payment.transaction}</div>
      <div class="field"><strong>Провайдер:</strong> ${order.payment.provider}</div>
      <div class="field"><strong>Сумма:</strong> ${order.payment.amount} ${order.payment.currency}</div>
      <div class="field"><strong>Стоимость доставки:</strong> ${order.payment.delivery_cost} ${order.payment.currency}</div>
      <div class="field"><strong>Общая стоимость товаров:</strong> ${order.payment.goods_total} ${order.payment.currency}</div>
      <div class="field"><strong>Банк:</strong> ${order.payment.bank}</div>
      <div class="field"><strong>Дата оплаты:</strong> ${formatTimestamp(order.payment.payment_dt)}</div>
    </div>
    <div class="section">
      <h3>Товары</h3>
      <table>
        <thead>
          <tr>
            <th>Название</th>
            <th>Бренд</th>
            <th>Цена</th>
            <th>Скидка (%)</th>
            <th>Итоговая цена</th>
            <th>Статус</th>
          </tr>
        </thead>
        <tbody>
          ${order.items.map(item => `
            <tr>
              <td>${item.name}</td>
              <td>${item.brand}</td>
              <td>${item.price}</td>
              <td>${item.sale}</td>
              <td>${item.total_price}</td>
              <td>${item.status}</td>
            </tr>
          `).join('')}
        </tbody>
      </table>
    </div>
    <div class="section">
      <small>Дата создания заказа: ${formatDate(order.date_created)}</small>
    </div>
  `;
}
