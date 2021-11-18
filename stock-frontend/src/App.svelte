<script>
	import Stock from '../../artifacts/contracts/Stock.sol/Stock.json'
	import { ethers } from 'ethers'
	import { onMount } from 'svelte'
	import detectEthereumProvider from '@metamask/detect-provider'



	export let name;
	let myStocks = 0, totalStocks = 0, nbrOwners = 0;
	const stockAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3"

	async function requestAccount() {
		await window.ethereum.request({ method: 'eth_requestAccounts' });
	}

	async function fetchStockQte() {
    if (typeof window.ethereum !== 'undefined') {
			await requestAccount()
      const provider = new ethers.providers.Web3Provider(window.ethereum)
			const signer = provider.getSigner()
      console.log({ provider })
      const contract = new ethers.Contract(stockAddress, Stock.abi, signer)
      try {
        const data = await contract.stockQte()
        myStocks = data.toNumber()
      } catch (err) {
        console.log("Error: ", err)
      }
    }
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

	async function fetchTotalSupply() {
		if (typeof window.ethereum !== 'undefined') {
			const provider = new ethers.providers.Web3Provider(window.ethereum)
			console.log({ provider })
			const contract = new ethers.Contract(stockAddress, Stock.abi, provider)
			try {
				const data = await contract.getTotalSupply()
				totalStocks = data.toNumber()
			} catch (err) {
				console.log("Error: ", err)
			}
		}
	}

	async function refresh() {
    if (typeof window.ethereum !== 'undefined') {
      try {
        await fetchStockQte()
				await fetchNbrOwners()
				await fetchTotalSupply()
      } catch (err) {
        console.log("Error: ", err)
      }
    }
  }

	onMount(async () => {
		refresh()
	});
</script>

<main>

	<div class="container">
		<div class="row justify-content-center">
			<h1>Decentralized {name}</h1>
		</div>
		<div class="row justify-content-center">
			<p>Dapp for decentralized stock-option</p>
		</div>

		<div class="row justify-content-center mt-5">
	    <div class="col mt-2">
				<div class="card" style="width: 18rem;">
				  <div class="card-header">
				    My amount of stocks
				  </div>
				  <ul class="list-group list-group-flush">
				    <li class="list-group-item">{myStocks}</li>
				  </ul>
				</div>
	    </div>
	    <div class="col mt-2">
				<div class="card" style="width: 18rem;">
				  <div class="card-header">
				    Total amount of stocks
				  </div>
				  <ul class="list-group list-group-flush">
				    <li class="list-group-item">{totalStocks}</li>
				  </ul>
				</div>
	    </div>
	    <div class="col mt-2">
				<div class="card" style="width: 18rem;">
				  <div class="card-header">
				    Number of stoks's owner
				  </div>
				  <ul class="list-group list-group-flush">
				    <li class="list-group-item">{nbrOwners}</li>
				  </ul>
				</div>
	    </div>
	  </div>

	</div>

</main>

<style>
	main {
		text-align: center;
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
