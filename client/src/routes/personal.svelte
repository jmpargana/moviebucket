<script>
	import { personalMovieBucket, dailyPick } from '../stores/movies';

	function handleAddMovie(e) {
		if (e.key === 'Enter') {
			personalMovieBucket.add({
				title: e.target.value
			});
			e.target.value = '';
		}
	}
</script>

<h2>Personal Area</h2>
<div class="row">
<h4>Daily Pick</h4>
<h1 class="daily-pick">{$dailyPick?.title}</h1>
</div>

<div class="input-block" on:keydown={handleAddMovie}>
	<label class="input-label" for="input">New Movie</label>
	<input id="input" />
</div>

<div class="movie-list">
	{#each $personalMovieBucket as movie}
		<div class="movie-row">{movie.title}</div>
	{/each}
</div>

<style>
	.input-block {
		--margin-top: 35px;
		margin-top: calc(var(--margin-top) + var(--space-xl));
		position: relative;
	}

	.input-label {
		font-family: var(--font-heading);
		position: absolute;
		top: calc(-1 * var(--margin-top));
		left: var(--space-xxs);
	}

.daily-pick {
  color: var(--color-primary);
}

.row {
  display: flex;
  align-items: flex-end;
  /* gap: var(--space-sm); */
}

.movie-list {
  margin-top: var(--space-xl);
}
</style>
