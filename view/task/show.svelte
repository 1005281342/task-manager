<script>
  export let task = {}
</script>

<h1>Show Task</h1>

<table border="1" cellpadding="10">
  <thead>
    {#each Object.keys(task) as key}
      <th>{key}</th>
    {/each}
  </thead>
  <tr>
    {#each Object.keys(task) as key}
      <td>{task[key]}</td>
    {/each}
  </tr>
</table>

<br />

<a href={`/task`}>Back</a>
<span> | </span>
<a href={`/task/${task.id || 0}/edit`}>Edit</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>
