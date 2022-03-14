<script>
  import UserSuggestion from './user-suggestion.svelte'

	export let title = '';

	const users = [
		{
			name: 'John Smith',
			movie: 'Harry Potter'
		},
		{
			name: 'John Smith',
			movie: 'Harry Potter'
		},
		{
			name: 'John Smith',
			movie: 'Harry Potter'
		},
		{
			name: 'John Smith',
			movie: 'Harry Potter'
		},
		{
			name: 'John Smith',
			movie: 'Harry Potter'
		},
		{
			name: 'John Smith',
			movie: 'Harry Potter'
		},
		{
			name: 'John Smith',
			movie: 'Harry Potter'
		},
		{
			name: 'John Smith',
			movie: 'Harry Potter'
		}
	];
</script>

<h3 class="movie-table-title">{title}</h3>
<div class="wrapper">
  {#each users as user}
    <UserSuggestion {...user} />
  {/each}
</div>

<style>
	.movie-table-title {
		color: var(--color-primary);
	}

.wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: var(--space-md);
  margin-top: var(--space-xs);
}

@media (min-width: 800px) {
  .wrapper {
    margin: var(--space-xl);
  }
}
</style>
