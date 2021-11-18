const hre = require("hardhat");

async function main() {

  const Stock = await ethers.getContractFactory("Stock");
  const stock = await Stock.deploy();
  await stock.deployed();

  console.log("Stock deployed to:", stock.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
