<script>
	import Favorite from '$lib/icons/favorite.svelte'
	import Add from '$lib/icons/add.svelte'
	const icons = {
		add: Add,
		favorite: Favorite,
	};
	export let icon;
  export let size = 30
  export let fill = "black"
</script>

<i class="icon-wrapper"><svelte:component width={size} height={size} this={icons[icon]} {fill} /></i>

<style>
	.icon-wrapper {
    width: 44px;
    height: 44px;
		display: grid;
		place-items: center;
		background-color: transparent;
	}

.icon-wrapper:hover {
  cursor: pointer;
  fill: var(--color-primary);
}
</style>


