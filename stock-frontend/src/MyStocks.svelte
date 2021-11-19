<script>
	import Stock from '../../artifacts/contracts/Stock.sol/Stock.json'
	import { ethers } from 'ethers'
	import { onMount } from 'svelte'

	export let myStocks;

</script>

<div class="card bg-primary">
  <div class="card-header">
    My amount of stocks
  </div>
  <ul class="list-group list-group-flush">
    <li class="list-group-item">{myStocks}</li>
  </ul>
</div>

<style>

</style>
