<script>
	import Stock from '../../artifacts/contracts/Stock.sol/Stock.json'
	import { ethers } from 'ethers'
	import { onMount } from 'svelte'

	let nbrOwners = 0;
	export let stockAddress;

	async function requestAccount() {
		await window.ethereum.request({ method: 'eth_requestAccounts' });
	}

	async function fetchNbrOwners() {
		if (typeof window.ethereum !== 'undefined') {
			const provider = new ethers.providers.Web3Provider(window.ethereum)
			console.log({ provider })
			const contract = new ethers.Contract(stockAddress, Stock.abi, provider)
			try {
				const data = await contract.nbrOwners()
				nbrOwners = data.toNumber()
			} catch (err) {
				console.log("Error: ", err)
			}
		}
	}

	onMount(async () => {
		fetchNbrOwners()
	});
</script>

<main>
	<div class="card" style="width: 18rem;">
	  <div class="card-header">
	    Number of stoks's owner
	  </div>
	  <ul class="list-group list-group-flush">
	    <li class="list-group-item">{nbrOwners}</li>
	  </ul>
	</div>
</main>

<style>

</style>
