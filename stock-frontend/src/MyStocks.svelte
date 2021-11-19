<script>
	import Stock from '../../artifacts/contracts/Stock.sol/Stock.json'
	import { ethers } from 'ethers'
	import { onMount } from 'svelte'

	let myStocks = 0;
	export let stockAddress;

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

	onMount(async () => {
		fetchStockQte()
	});
</script>

<div class="card">
  <div class="card-header">
    My amount of stocks
  </div>
  <ul class="list-group list-group-flush">
    <li class="list-group-item">{myStocks}</li>
  </ul>
</div>

<style>

</style>
