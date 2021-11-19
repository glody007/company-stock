<script>
	import Stock from '../../artifacts/contracts/Stock.sol/Stock.json'
	import { ethers } from 'ethers'
	import { onMount } from 'svelte'

	export let totalStocks = 0;

</script>

<div class="card bg-info">
  <div class="card-header">
    Total amount of stocks
  </div>
  <ul class="list-group list-group-flush">
    <li class="list-group-item">{totalStocks}</li>
  </ul>
</div>

<style>

</style>
