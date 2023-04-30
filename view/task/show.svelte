<script>
  export let task = {}
</script>

<div class="container">
  <h1>Show Task</h1>

  <div class="form">
    <div class="form-group">
      <p><strong>ID:</strong></p>
      <p>{task.ID}</p>
    </div>

    <div class="form-group">
      <p><strong>TaskType:</strong></p>
      <p>{task.task_type}</p>
    </div>

    <div class="form-group">
      <p><strong>Payload:</strong></p>
      <p>{task.payload}</p>
    </div>

    <div class="form-group">
      <p><strong>Cron Expression:</strong></p>
      <p>{task.cron_spec}</p>
    </div>
  </div>

  <div class="button-group">
    <a href="/task" class="btn-back">Back</a>
    <a href={`/task/${task.ID || 0}/edit`} class="btn-edit">Edit</a>
  </div>
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
    margin-top: 20px;
    width: 100%;
    max-width: 500px;
  }

  /* 修改 .btn-back 按钮和 .btn-edit 按钮的样式，增加 padding 和 margin，使其看起来更为突出 */
  .btn-back {
    color: #333;
    display: inline-block;
    font-size: 0.9rem;
    text-decoration: none;
    text-transform: uppercase;
    transition: color 0.3s ease;
    padding: 10px 15px;
    margin: 0;
    border-radius: 4px;
    border: 1px solid #333;
  }

  .btn-edit {
    color: #fff;
    display: inline-block;
    font-size: 0.9rem;
    text-decoration: none;
    text-transform: uppercase;
    transition: color 0.3s ease;
    padding: 10px 15px;
    margin: 0;
    border-radius: 4px;
    background-color: green;
  }

  /* 修改 .btn-back 按钮和 .btn-edit 按钮的 hover 样式，使其看起来更为突出 */
  .btn-back:hover {
    color: #fff;
    background-color: #333;
  }

  .btn-edit:hover {
    background-color: darkgreen;
  }

  /* 修改 .form-group 的样式，使其与 view/task/new.svelte 中的表单样式保持一致 */
  .form-group {
    margin-bottom: 20px;
    border-top: 1px solid lightgray;
    padding: 10px;
    border-radius: 4px;
  }

  .form-group:first-of-type {
    border-top: none;
  }

  strong {
    font-weight: bold;
  }

  p {
    overflow-wrap: break-word;
    word-wrap: break-word;
    word-break: break-all;
  }
</style>
