<script>
  export let tasks = []
  const keys = Object.keys(tasks[0] || {})
</script>

<div class="container">
  <h1>Task Index</h1>

  <table role="table" aria-label="List of tasks">
    <thead>
    {#if tasks.length > 0}
      <tr>
        {#each keys as key}
          {#if key.toLowerCase() === "id"}
            <th class="id">{key}</th>
          {:else}
            <th>{key}</th>
          {/if}
        {/each}
      </tr>
    {/if}
    </thead>

    <tbody>
    {#each tasks as task}
      <tr>
        {#each keys as key}
          {#if key.toLowerCase() === 'id'}
            <td class="id"><a href={`/task/${task[key] || 0}`}>{task[key]}</a></td>
          {:else}
            <td>{task[key]}</td>
          {/if}
        {/each}
      </tr>
    {/each}
    </tbody>
  </table>

  <button class="btn-new"><a href={`/task/new`}>Create Task</a></button>
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

  .container {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  h1 {
    font-weight: 600;
    margin: 20px 0;
    text-align: center;
  }

  table {
    background-color: #fff;
    border-collapse: collapse;
    border-radius: 10px;
    box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
    margin: 20px 0;
    overflow: hidden;
    width: 100%;
  }

  th,
  td {
    padding: 12px;
    text-align: left;
  }

  td {
    font-size: 0.9rem;
  }

  th {
    background-color: #f2f2f2;
    font-size: 0.95rem;
    text-transform: uppercase;
  }

  tbody tr:hover {
    background-color: #f7f7f7;
  }

  a {
    color: #333;
    text-decoration: none;
  }

  .btn-new {
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

  .btn-new:hover {
    background-color: #3e8e41;
  }

  /* 用户ID链接样式 */
  td.id a,
  th.id {
    color: #2196f3;
    text-decoration: none;
  }

  td.id a:hover {
    text-decoration: underline;
  }
</style>
