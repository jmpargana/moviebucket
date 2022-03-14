<script>
  import '../app.css'
  import Navbar from '$lib/navbar.svelte'
</script>

<Navbar />
<main class="main-block">
  <slot />
</main>

<style>
.main-block {
  padding: var(--space-md);
}
</style>
