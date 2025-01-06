-- https://leetcode.com/problems/product-sales-analysis-i/description

select product.product_name, sales.year, sales.price from sales
inner join product
on sales.product_id = product.product_id