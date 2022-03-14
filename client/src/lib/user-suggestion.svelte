<script>
  import Icon from './icon.svelte'

  export let name
  export let movie
</script>

<div class="user-suggestion-wrapper">
  <div class="top-row">
    <h5 class="user">{name}</h5>
    <Icon icon="favorite" />
  </div>
  <div class="bottom-row">
    <h5>{movie}</h5>
    <Icon icon="add" />
  </div>
</div>

<style>
.user-suggestion-wrapper {
  max-width: 800px;
  width: 100%;
  height: 200px;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  padding: var(--space-md);
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
  border-radius: var(--space-xs);
}


.top-row {
  display: flex;
  gap: var(--space-sm);
  align-items: center;
}

.bottom-row {
  display: flex;
  align-items: center;
  gap: var(--space-md);
  justify-content: center;
}

.user {
  color: var(--color-secondary);
}
</style>
