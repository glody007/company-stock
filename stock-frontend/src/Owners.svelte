<script>
	import Stock from '../../artifacts/contracts/Stock.sol/Stock.json'
	import { ethers } from 'ethers'
	import { onMount } from 'svelte'

  export let stockAddress;
  let owners = []

	async function requestAccount() {
		await window.ethereum.request({ method: 'eth_requestAccounts' });
	}

  async function fetchOwners() {
    if (typeof window.ethereum !== 'undefined') {
      const provider = new ethers.providers.Web3Provider(window.ethereum)
      console.log({ provider })
      const contract = new ethers.Contract(stockAddress, Stock.abi, provider)
      try {
        const data = await contract.getOwners()
        owners = data.toString().split(',')
      } catch (err) {
        console.log("Error: ", err)
      }
    }
	}

  async function getStocksOf(address) {
    if (typeof window.ethereum !== 'undefined') {
      const provider = new ethers.providers.Web3Provider(window.ethereum)
      console.log({ provider })
      const contract = new ethers.Contract(stockAddress, Stock.abi, provider)
      try {
        const data = await contract.stockQteOf(address)
        return data.toNumber()
      } catch (err) {
        console.log("Error: ", err)
      }
    }
  }

	onMount(async () => {
		fetchOwners()
	});
</script>

<div class="card">
  <table class="table">
  <thead class="thead-dark">
    <tr>
      <th scope="col">#</th>
      <th scope="col">Addresse</th>
      <th scope="col">Stocks</th>
    </tr>
  </thead>
  <tbody>
    {#each owners as owner, i}
    <tr>
      <th scope="row">{i}</th>
      <td>{owner}</td>
      <td>
        {#await getStocksOf(owner)}
        	<p>...waiting</p>
        {:then data}
        	<p>{data}</p>
        {:catch error}
        	<p style="color: red">{error.message}</p>
        {/await}
      </td>
    </tr>
	  {/each}
  </tbody>
</table>
</div>

<style>

</style>
