<div class="container">
  <h1>New Task</h1>

  <form method="post" action="/task" class="form">
    <div class="form-group">
      <label for="task_type">TaskType:</label>
      <input type="text" id="task_type" name="task_type" required>
    </div>

    <div class="form-group">
      <label for="payload">Payload:</label>
      <textarea id="payload" name="payload" required></textarea>
    </div>

    <div class="form-group">
      <label for="cron_spec">Cron Expression:</label>
      <input type="text" id="cron_spec" name="cron_spec" required>
    </div>

    <div class="button-group">
      <a href="/task" class="btn-back">Back</a>
      <button type="submit" class="btn-create">Create Task</button>
    </div>
  </form>
</div>

<style>
  *,
  *::before,
  *::after {
    box-sizing: border-box;
  }

  body {
    background-color: #f2f2f2;
    font-family: 'Poppins', sans-serif;
    margin: 0;
    padding: 0;
  }

  h1 {
    font-weight: 600;
    margin: 20px 0;
    text-align: center;
  }

  form.form {
    background-color: #fff;
    border-radius: 10px;
    box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
    padding: 20px;
    margin: 20px 0;
    overflow: hidden;
    width: 100%;
    max-width: 500px;
  }

  .form-group {
    margin-bottom: 20px;
  }

  label {
    display: block;
    font-size: 0.9rem;
    margin-bottom: 5px;
    text-transform: uppercase;
  }

  input[type="text"],
  textarea {
    background-color: #f2f2f2;
    border: none;
    border-radius: 4px;
    box-shadow: inset 0px 2px 1px rgba(255, 255, 255, 0.2);
    padding: 12px 15px;
    width: 100%;
  }

  input[type="text"]:focus,
  textarea:focus {
    outline: none;
    box-shadow: 0 0 2px #4CAF50;
  }

  input[type="submit"],
  button.btn-create {
    background-color: #4CAF50;
    border: none;
    border-radius: 4px;
    color: #fff;
    cursor: pointer;
    font-weight: 600;
    padding: 10px 20px;
    text-transform: uppercase;
    transition: background-color 0.3s ease;
  }

  input[type="submit"]:hover,
  button.btn-create:hover {
    background-color: #3e8e41;
  }

  .container {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  /* 添加以下样式，将 .form 和 .button-group 容器都居中 */
  .form,
  .button-group {
    margin: 0 auto;
  }

  /* 添加以下样式，使用 flex 布局将 .button-group 容器的子元素水平平分容器宽度 */
  .button-group {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    max-width: 500px;
  }

  /* 修改 .btn-back 按钮的样式，增加 padding 和 margin，使其看起来更为突出 */
  .btn-back {
    color: #333;
    display: inline-block;
    font-size: 0.9rem;
    text-decoration: none;
    text-transform: uppercase;
    transition: color 0.3s ease;
    padding: 10px 15px;
    margin-right: 10px;
    border-radius: 4px;
    border: 1px solid #333;
  }

  /* 修改 .btn-back 按钮的 hover 样式，使其看起来更为突出 */
  .btn-back:hover {
    color: #fff;
    background-color: #333;
  }
</style>
