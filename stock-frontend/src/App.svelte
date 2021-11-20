<script>
	import Content from './Content.svelte';
	import Modal from 'svelte-simple-modal';
	export let name, stockAddress;

</script>

<main>
	<Modal>
		<Content name={name} stockAddress={stockAddress}  />
	</Modal>
</main>

<style>
	main {
		background-image: url("background.jpg");
	}
</style>
