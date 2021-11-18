<script>
	import Greeter from '../../artifacts/contracts/Stock.sol/Stock.json'
	export let name;
</script>

<main>
	<h1>Decentralized {name}</h1>
	<p>Dapp for decentralized stock-option</p>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #5a57fa;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
