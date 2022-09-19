module Lib
    ( ant
    ) where

import Data.List ( group, concat )
import Control.Monad ( (>=>) )

a = sequence [ head ]


antList :: [String]
antList = iterate (group >=> concat . sequence [\t -> [head t], show . length]) "1"
antNumList :: [[Int]]
antNumList = iterate (group >=> sequence [head, length]) [1]

ant :: IO ()
-- ant = putStrLn $ foldr (\a b -> show a ++ b) "" (antList !! 100)
ant = print $ antList !! 10000 !! 10000

half x = if even x
  then Just $ x `div` 2
  else Nothing
