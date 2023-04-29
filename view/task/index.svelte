<script>
  export let tasks = []
</script>

<h1>Task Index</h1>

<table border="1" cellpadding="10">
  {#if tasks.length > 0}
    <thead>
      {#each Object.keys(tasks[0]) as key}
        <th>{key}</th>
      {/each}
    </thead>
  {/if}
  {#each tasks as task}
    <tr>
      {#each Object.keys(task) as key}
        {#if key.toLowerCase() === "id"}
          <td><a href={`/task/${task.id || 0}`}>{task[key]}</a></td>
        {:else}
          <td>{task[key]}</td>
        {/if}
      {/each}
    </tr>
  {/each}
</table>

<br />

<a href={`/task/new`}>New Task</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>
