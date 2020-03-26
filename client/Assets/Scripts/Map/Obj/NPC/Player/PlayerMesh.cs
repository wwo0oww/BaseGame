using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace NMObj
{
    public partial class Player : NPC
    {

        public override int OffY() {
            return 1;
        }


        public override void GenerateMeshData()
        {

            int x = 0, y = 0, z = 0;
            byte brick = 1;
            // Left wall
            BuildFace(brick, new Vector3(x, y, z), Vector3.up, Vector3.forward, false);
            // Right wall
            BuildFace(brick, new Vector3(x + 1, y, z), Vector3.up, Vector3.forward, true);

            // Bottom wall
            BuildFace(brick, new Vector3(x, y, z), Vector3.forward, Vector3.right, false);
            // Top wall
            BuildFace(brick, new Vector3(x, y + 1, z), Vector3.forward, Vector3.right, true);

            // Back
            BuildFace(brick, new Vector3(x, y, z), Vector3.up, Vector3.right, true);
            // Front
            BuildFace(brick, new Vector3(x, y, z + 1), Vector3.up, Vector3.right, false);



        }

        public override Rect GetTextureAtlasUv()
        {
            return Map.Instance.TextureAtlasUvs[6];
        }

        //public override void SetTextureAtlasUv()
        //{
        //    var brick = 3;
        //    Vector2 uvWidth = new Vector2(0.25f, 0.25f);
        //    Vector2 uvCorner = new Vector2(0.00f, 0.75f);
        //    uvCorner.x += (float)(brick - 1) / 4;
        //    Uvs.Add(uvCorner);
        //    Uvs.Add(new Vector2(uvCorner.x, uvCorner.y + uvWidth.y));
        //    Uvs.Add(new Vector2(uvCorner.x + uvWidth.x, uvCorner.y + uvWidth.y));
        //    Uvs.Add(new Vector2(uvCorner.x + uvWidth.x, uvCorner.y));
        //}
    }
}
