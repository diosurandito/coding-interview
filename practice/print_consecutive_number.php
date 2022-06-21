<?php
// Optimized PHP program to find
// sequences of all consecutive
// numbers with sum equal to N
 
function printSums($N)
{
    $start = 1; $end = 1;
    $sum = 1;
 
    while ($start <= $N / 2)
    {
        if ($sum < $N)
        {
            $end += 1;
            $sum += $end;
        }
        else if ($sum > $N)
        {
            $sum -= $start;
            $start += 1;
        }
        else if ($sum == $N)
        {
            for ($i = $start; $i <= $end; ++$i){
                echo $i," ";
                // echo "\n";
            }
            
            $sum -= $start;
            $start += 1;
            
        }
    }
}
 
// Driver Code
    printSums(6);
?>