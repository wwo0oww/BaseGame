using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace NMObj
{
    public class ObjAction : MonoBehaviour
    {
        private float baseMove = 1f;
        private Obj m_obj;
        private int m_nMoveFrameCount;
        private int m_scale;
        private uint m_count = 0;

        //主摄像机对象
        public Camera Camera;
        //NPC名称
        private string m_name = "";

        //主角对象
        GameObject hero;
        //NPC模型高度
        float npcHeight;
        //红色血条贴图
        public Texture2D blood_red;
        //黑色血条贴图
        public Texture2D blood_black;
        //默认NPC血值
        private int HP = 100;


        public Obj ObjNode
        {
            get { return m_obj; }
            set { m_obj = value; }
        }
        // Use this for initialization
        void Start()
        {
            if (m_obj != null && m_obj.ObjType == T_Obj.T_TERRAIN.GenObjType())
            {
                return;
            }
            m_nMoveFrameCount = 0;
            m_scale = Game.nFrame / Game.nServerFrame;
            baseMove = baseMove * Config.Scale / m_scale;

        }


        void FixedUpdate()
        {
            if (m_obj != null && m_obj.ObjType == T_Obj.T_TERRAIN.GenObjType())
            {
                return;
            }
            if (m_obj != null)
            {

                //
                if (Map.Instance != null && m_obj.ID == Map.Instance.Player.ID)
                {
                    if (Map.Instance != null && Map.Instance.ServerFrame != 0)
                        m_count++;
                    if (m_count % m_scale == 0)
                    {
                        if (Map.Instance != null && Map.Instance.ServerFrame != 0)
                        {
                            //Debug.Log(Map.Instance.ServerFrame);
                            Map.Instance.ServerFrame++;
                        }
                    }
                }
                if ((m_obj.GPose.x != gameObject.transform.position.x ||
                    m_obj.GPose.y != gameObject.transform.position.y))
                {
                    gameObject.transform.position = new Vector3(m_obj.GPose.x, gameObject.transform.position.y, m_obj.GPose.z);
                }
                if (m_obj.NextDirection != ObjDirection.NONE)
                {

                    if (m_nMoveFrameCount % m_scale == 0)
                    {
                        m_nMoveFrameCount = 0;
                        if (m_obj.NextDirection == ObjDirection.STOP)
                        {
                            m_obj.DoStopMove();
                        }
                        else
                        {
                            m_obj.DoMove(m_obj.NextDirection);
                        }
                    }

                }
                Debug.Log(m_obj.Name + m_obj.HasStatus(ObjStatus.MOVE)+ m_obj.Direction);
                if (m_obj.HasStatus(ObjStatus.MOVE))
                {
                    move();
                }
            }
            CheckRecycle();
        }

        void CheckRecycle()
        {
            if (m_obj == null)
            {
                return;
            }
            if (Mathf.Abs(Map.Instance.MapCenterPos.x - m_obj.GPose.x) > 200 ||
                Mathf.Abs(Map.Instance.MapCenterPos.y - m_obj.GPose.y) > 200)
            {
                ObjMgr.RecycleObj(m_obj);
            }
        }

        private void move()
        {
            if (m_obj.ID == Map.Instance.Player.ID && m_count % m_scale == 0)
            {
                //Debug.Log(gameObject.transform.position);
            }
            Vector3 OffPos;
            switch (m_obj.Direction)
            {
                case ObjDirection.LEFT:
                    OffPos = new Vector3(-baseMove, 0, 0);
                    gameObject.transform.position += OffPos;
                    m_obj.GPose += OffPos;
                    m_nMoveFrameCount++;
                    break;
                case ObjDirection.RIGHT:
                    OffPos = new Vector3(baseMove, 0, 0);
                    gameObject.transform.position += OffPos;
                    m_obj.GPose += OffPos;
                    m_nMoveFrameCount++;
                    break;
                case ObjDirection.FORWARD:
                    OffPos = new Vector3(0, 0, baseMove);
                    gameObject.transform.position += OffPos;
                    m_obj.GPose += OffPos;
                    m_nMoveFrameCount++;
                    break;
                case ObjDirection.BACK:
                    OffPos = new Vector3(0, 0, -baseMove);
                    gameObject.transform.position += OffPos;
                    m_obj.GPose += OffPos;
                    m_nMoveFrameCount++;
                    break;
            }

        }

        void OnGUI()
        {
            if (m_obj != null && m_obj.ObjType == T_Obj.T_TERRAIN.GenObjType())
            {
                return;
            }
            if (m_obj != null && m_name == "")
            {
                m_name = m_obj.Name;
                if (m_obj.ObjType == T_Obj.T_NPC.T_PLAYER.GenObjType())
                {
                    m_name = "Player:" + m_name;
                    GUI.color = Color.red;
                }
                else
                {
                    m_name = "NPC:" + m_name;
                    GUI.color = Color.white;
                }
            }
            //得到NPC头顶在3D世界中的坐标
            //默认NPC坐标点在脚底下，所以这里加上npcHeight它模型的高度即可
            var npcHeight = 1;
            Vector3 worldPosition = new Vector3(transform.position.x + 0.2f, transform.position.y + npcHeight, transform.position.z);
            //根据NPC头顶的3D坐标换算成它在2D屏幕中的坐标
            Vector2 position = Camera.WorldToScreenPoint(worldPosition);
            //得到真实NPC头顶的2D坐标
            position = new Vector2(position.x, Screen.height - position.y);
            //注解2
            //计算出血条的宽高
            // Vector2 bloodSize = GUI.skin.label.CalcSize(new GUIContent(blood_red));

            //通过血值计算红色血条显示区域
            // int blood_width = blood_red.width * HP / 100;
            //先绘制黑色血条
            //GUI.DrawTexture(new Rect(position.x - (bloodSize.x / 2), position.y - bloodSize.y, bloodSize.x, bloodSize.y), blood_black);
            //在绘制红色血条
            // GUI.DrawTexture(new Rect(position.x - (bloodSize.x / 2), position.y - bloodSize.y, blood_width, bloodSize.y), blood_red);

            //注解3
            //计算NPC名称的宽高
            Vector2 nameSize = GUI.skin.label.CalcSize(new GUIContent(m_name));
            //设置显示颜色为黄色


            //绘制NPC名称
            GUI.Label(new Rect(position.x - (nameSize.x / 2), position.y - nameSize.y, nameSize.x, nameSize.y), m_name);

        }
    }

}