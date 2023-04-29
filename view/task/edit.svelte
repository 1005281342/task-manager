<script>
  export let task = {}
</script>

<h1>Edit Task</h1>

<form method="post" action={`/task/${task.id || 0}`}>
  <input type="hidden" name="_method" value="patch" />
  <!-- Add input fields here -->
  <input type="submit" value="Update Task" />
</form>

<br />

<a href={`/task`}>Back</a>
<span> | </span>
<a href={`/task/${task.id || 0}`}>Show Task</a>
