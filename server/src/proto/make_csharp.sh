cd proto/net
protoc --csharp_out=../../net *.proto
cd ../rpc
protoc --csharp_out=../../rpc *.proto
cp ../../net/*.cs E:/code/mine/MinePackage-master/Assets/Scripts/Proto