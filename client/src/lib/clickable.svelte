<script>
  import Icon from '$lib/icon.svelte'

  export let value
  export let type
</script>

<div class="clickable-wrapper">
  <span>{value}</span>
  {#if type === 'follow'}
    <Icon icon="favorite" />
  {:else}
    <Icon icon="add" />
  {/if}
</div>
