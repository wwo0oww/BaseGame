using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
namespace NMObj
{
    public enum ObjStatus
    {
        NONE = 1,
        MOVE = 2,
        ATTACK = 4,
    }
    public enum ObjDirection
    {
        NONE = 0,
        LEFT = 1,
        RIGHT = 2,
        FORWARD = 3,
        BACK = 4,
        STOP = 5,
    }
    public partial class Obj
    {
        private Vector3 m_pos;
        private int m_objType;
        private Transform m_objTransform;
        private int m_id;
        private ObjStatus m_status = ObjStatus.NONE;

        private ObjDirection m_direction = ObjDirection.NONE;

        private ObjDirection m_nextdirection = ObjDirection.NONE;

        private ObjAction m_objAction;

        private int m_speed;

        private string m_name;

        private long m_lastMoveProto = 0;

        public string Name
        {
            get { return m_name; }
        }

        public int ID
        {
            get { return m_id; }
        }

        public int Speed
        {
            get { return m_speed; }
        }

        public ObjAction ObjAction
        {
            get { return m_objAction; }
            set { m_objAction = value; }
        }

        public ObjDirection Direction
        {
            get { return m_direction; }
        }

        public ObjDirection NextDirection
        {
            get { return m_nextdirection; }
        }

        public ObjStatus Status
        {
            get { return m_status; }
        }
        public Vector3 GPose
        {
            get { return m_pos; }
            set
            {
                m_pos = new Vector3(value.x, m_pos.y, value.z);
            }
        }
        public int ObjType
        {
            get { return m_objType; }
            set { m_objType = value; }
        }
        public Transform ObjTransform
        {
            get { return m_objTransform; }
            set { m_objTransform = value; }
        }
        public Obj(Vector3 pos, int objType, int id)
        {
            this.m_pos = pos;
            this.ObjType = objType;
            this.m_id = id;
        }
        public virtual void Init()
        {
            if (ObjTransform == null)
            {
                ObjTransform = UnityEngine.Object.Instantiate(Map.Instance.ObjPrefab, this.GPose, Quaternion.identity);
                var mainTexture = GetMainTexture();
                if (mainTexture != null)
                    ObjTransform.GetComponent<Renderer>().sharedMaterial.mainTexture = mainTexture;
                ObjTransform.parent = Map.Instance.ObjParent;
                ObjTransform.name = ToString();

                ObjAction = ObjTransform.GetComponent<ObjAction>();
                ObjAction.ObjNode = this;

                InitMeshData();
            }
        }

        public void SetID(int ID)
        {
            m_id = ID;
        }

        public void SetSpeed(int speed)
        {
            m_speed = speed;
        }

        public void SetName(string name)
        {
            m_name = name;
        }

        public override string ToString()
        {
            return string.Format("{0}_{1}", m_objType, m_id);
        }

        public void Destroy()
        {
            if (m_objTransform != null)
            {
                MeshFilter meshFilter = m_objTransform.GetComponent<MeshFilter>();
                meshFilter.mesh.Clear();
                UnityEngine.Object.Destroy(meshFilter.mesh);
                meshFilter.mesh = null;
                UnityEngine.Object.Destroy(meshFilter);
                UnityEngine.Object.Destroy(m_objTransform);
                m_objTransform = null;

                Vertices.Clear();
                Uvs.Clear();
                Colors.Clear();
                Indices.Clear();
            }
        }
        public bool HasStatus(ObjStatus status)
        {
            return (Status & status) != 0;
        }

        public void AddStatus(ObjStatus status)
        {
            m_status |= status;
        }

        public bool IsMove()
        {
            return Status == ObjStatus.MOVE && Direction != ObjDirection.NONE;
        }

        public void Move(ObjDirection direction)
        {
            if (Map.Instance == null || (Map.Instance.Player!=null&&ID != Map.Instance.Player.ID))
            {
                DoMove(direction);
            }
            if (HasStatus(ObjStatus.MOVE) && direction == m_direction)
            {
                return;
            }
            if (Game.SinglePlayer == 0)
            {
                if (DateTime.Now.Ticks - m_lastMoveProto > 200 * 10000)
                {
                    m_lastMoveProto = DateTime.Now.Ticks;
                    Net.m_obj_move_tos t1 = new Net.m_obj_move_tos();
                    t1.Direction = (int)direction;
                    WSocket.Send(t1);
                }

            }
            else
            {
                m_nextdirection = direction;
            }
        }

        public void DoMove(ObjDirection direction)
        {
            m_status = ObjStatus.MOVE;
            m_direction = direction;
            m_nextdirection = ObjDirection.NONE;
        }

        public void StopMove()
        {
            if (!HasStatus(ObjStatus.MOVE))
            {
                return;
            }
            if (Map.Instance == null || ID != Map.Instance.Player.ID)
            {
                DoStopMove();
            }
            if (Game.SinglePlayer == 0)
            {
                if (DateTime.Now.Ticks - m_lastMoveProto > 200 * 10000)
                {
                    m_lastMoveProto = DateTime.Now.Ticks;
                    Net.m_obj_move_tos t1 = new Net.m_obj_move_tos();
                    t1.Direction = (int)(ObjDirection.STOP);
                    WSocket.Send(t1);
                }
            }
            else
            {
                m_nextdirection = ObjDirection.STOP;
            }
        }
        public void DoStopMove()
        {
            m_status = ObjStatus.NONE;
            m_direction = ObjDirection.NONE;
            m_nextdirection = ObjDirection.NONE;
        }

    }
}
